package fanfou

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mogita/oauth"
)

var (
	// LibraryVersion represents this library version
	LibraryVersion = "0.1"

	// UserAgent represents this client User-Agent
	UserAgent = "github.com/mogita/go-fanfou v" + LibraryVersion

	// BaseURL represents Fanfou API base URL
	BaseURL = "http://api.fanfou.com/"

	// AuthBaseURL represents Fanfou API authorization base URL
	AuthBaseURL = "http://fanfou.com/"

	// Request token URI
	requestTokenURI = "oauth/request_token"

	// Authorize token URI
	authorizeTokenURI = "oauth/authorize"

	// Access token URI
	accessTokenURI = "oauth/access_token"
)

// A Client manages communication with the Fanfou API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// OAuth consumer used to handle authentication work
	oauthConsumer *oauth.Consumer

	// Base URL for API requests.
	BaseURL *url.URL

	// UserAgent agent used when communicating with Fanfou API.
	UserAgent string

	// Application consumer key
	ConsumerKey string

	// Application consumer secret
	ConsumerSecret string

	// Services used for talking to different parts of the API.
	Users         *UsersService
	Statuses      *StatusesService
	Search        *SearchService
	Trends        *TrendsService
	Blocks        *BlocksService
	Account       *AccountService
	SavedSearches *SavedSearchesService
	Photos        *PhotosService
	Followers     *FollowersService
	Favorites     *FavoritesService

	// Temporary Response
	Response *Response
}

// NewClient returns a new Fanfou API client. if a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(consumerKey, consumerSecret string) *Client {
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		BaseURL:        baseURL,
		UserAgent:      UserAgent,
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		oauthConsumer: oauth.NewConsumer(
			consumerKey,
			consumerSecret,
			oauth.ServiceProvider{
				RequestTokenUrl:   AuthBaseURL + requestTokenURI,
				AuthorizeTokenUrl: AuthBaseURL + authorizeTokenURI,
				AccessTokenUrl:    AuthBaseURL + accessTokenURI,
			},
		),
	}

	c.oauthConsumer.Debug(false)

	c.Users = &UsersService{client: c}
	c.Statuses = &StatusesService{client: c}
	c.Search = &SearchService{client: c}
	c.Trends = &TrendsService{client: c}
	c.Blocks = &BlocksService{client: c}
	c.Account = &AccountService{client: c}
	c.SavedSearches = &SavedSearchesService{client: c}
	c.Photos = &PhotosService{client: c}
	c.Followers = &FollowersService{client: c}
	c.Favorites = &FavoritesService{client: c}

	return c
}

// GetRequestTokenAndURL returns the request token and the login url for authorizing this token
func (c *Client) GetRequestTokenAndURL(callbackURL string) (*oauth.RequestToken, string, error) {
	rToken, loginURL, err := c.oauthConsumer.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		return nil, "", err
	}

	return rToken, loginURL, nil
}

// AuthorizeClient completes the OAuth authorization to the client
// so it can communicate with Fanfou API
func (c *Client) AuthorizeClient(rToken *oauth.RequestToken) error {
	accessToken, err := c.oauthConsumer.AuthorizeToken(rToken, "")

	if err != nil {
		errResp := new(ErrorResponse)
		errResp.Meta = &ResponseMeta{
			Error:   "unknown authorization error",
			Request: "",
		}

		if oauthErr, ok := err.(oauth.HTTPExecuteError); ok {
			// fanfou auth error body is in XML
			errResp.Response = oauthErr.Response
			errXML := xml.Unmarshal(oauthErr.ResponseBodyBytes, &errResp.Meta)
			if errXML != nil {
				errResp.Meta.Error = errXML.Error()
				return errResp
			}

			return errResp
		}

		errResp.Meta.Error = err.Error()
		return err
	}

	c.client, err = c.oauthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}

// AuthorizeClientWithXAuth completes the OAuth authorization to the client
// with XAuth so it can communicate with Fanfou API
//
// This method is a simplified OAuth process, taking username and password
// to authorize the client, without the need to redirect to the web UI
func (c *Client) AuthorizeClientWithXAuth(username, password string) error {
	c.oauthConsumer.AdditionalParams["x_auth_username"] = username
	c.oauthConsumer.AdditionalParams["x_auth_password"] = password
	c.oauthConsumer.AdditionalParams["x_auth_mode"] = "client_auth"

	reqToken := oauth.RequestToken{}
	accessToken, err := c.oauthConsumer.AuthorizeToken(&reqToken, "")

	if err != nil {
		errResp := new(ErrorResponse)
		errResp.Meta = &ResponseMeta{
			Error:   "unknown authorization error",
			Request: "",
		}

		if oauthErr, ok := err.(oauth.HTTPExecuteError); ok {
			// fanfou auth error body is in XML
			errResp.Response = oauthErr.Response
			errXML := xml.Unmarshal(oauthErr.ResponseBodyBytes, &errResp.Meta)
			if errXML != nil {
				errResp.Meta.Error = errXML.Error()
				return errResp
			}

			return errResp
		}

		errResp.Meta.Error = err.Error()
		return err
	}

	c.client, err = c.oauthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}

// NewRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(method, uri string, body string) (*http.Request, error) {
	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %+v", err)
		}
	}()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	response := new(Response)
	if v != nil {
		response.Data = v
		err = json.NewDecoder(resp.Body).Decode(response.Data)
		c.Response = response
	}

	return resp, err
}

// Response specifies Fanfou's response structure.
type Response struct {
	Response *http.Response // HTTP response
	Data     interface{}
	Meta     *ResponseMeta
}

// ResponseMeta represents information about the response. If all goes well,
// only a Code key with value 200 will present. However, sometimes things
// go wrong, and in that case ErrorType and ErrorMessage are present.
type ResponseMeta struct {
	Request string `json:"request,omitempty" xml:"request"`
	Error   string `json:"error,omitempty" xml:"error"`
}

// ErrorResponse represents a Response which contains an error
type ErrorResponse Response

// Error implements the error interface
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Meta.Error)
}

// GetStatusCode gets the status code of the error response
func (r *ErrorResponse) GetStatusCode() string {
	return fmt.Sprintf("%d", r.Response.StatusCode)
}

// GetRequestMethod gets the request method of the error response
func (r *ErrorResponse) GetRequestMethod() string {
	return fmt.Sprintf("%s", r.Response.Request.Method)
}

// GetRequestURL gets the request url of the error response
func (r *ErrorResponse) GetRequestURL() string {
	return fmt.Sprintf("%s", r.Response.Request.URL)
}

// GetFanfouError gets the error message returned by Fanfou API
// if presented in the response
func (r *ErrorResponse) GetFanfouError() string {
	return fmt.Sprintf("%s", r.Meta.Error)
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(res *http.Response) error {
	if res.StatusCode == http.StatusOK {
		return nil
	}

	r := new(ErrorResponse)
	r.Response = res
	// default error message
	r.Meta = &ResponseMeta{
		Error:   "unknown error",
		Request: res.Request.URL.String(),
	}

	if res.StatusCode >= http.StatusInternalServerError {
		r.Meta.Error = http.StatusText(res.StatusCode)
		return r
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.Meta.Error = err.Error()
	}

	if err := json.Unmarshal(data, &r.Meta); err != nil {
		r.Meta.Error = err.Error()
	}

	return r
}
