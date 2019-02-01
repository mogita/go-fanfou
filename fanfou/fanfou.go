package fanfou

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mogita/oauth"
)

var (
	// LibraryVersion represents this library version
	LibraryVersion = "1"

	// UserAgent represents this client User-Agent
	UserAgent = "github.com/mogita/go-fanfou v" + LibraryVersion

	// BaseURL represents Fanfou API base URL
	BaseURL = "https://api.fanfou.com/"

	// AuthBaseURL represents Fanfou API authorization base URL
	AuthBaseURL = "https://fanfou.com/"

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
	Users          *UsersService
	Statuses       *StatusesService
	Search         *SearchService
	Trends         *TrendsService
	Blocks         *BlocksService
	Account        *AccountService
	SavedSearches  *SavedSearchesService
	Photos         *PhotosService
	Followers      *FollowersService
	Favorites      *FavoritesService
	Friends        *FriendsService
	Friendships    *FriendshipsService
	DirectMessages *DirectMessagesService

	// Temporary Response
	Response *Response
}

// ReqeustToken provides the structure as oauth.RequestToken
type RequestToken struct {
	Token  string
	Secret string
}

// AccessToken provides the structure as oauth.AccessToken
type AccessToken struct {
	Token          string
	Secret         string
	AdditionalData map[string]string
}

// NewClient returns a new Fanfou API client.
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
	c.Friends = &FriendsService{client: c}
	c.Friendships = &FriendshipsService{client: c}
	c.DirectMessages = &DirectMessagesService{client: c}

	return c
}

// GetRequestTokenAndURL returns the request token and the login url for authorizing this token.
//
// "callbackURL" can be "oob" if you're running your application outside a browser.
//
// Read more about "oob" authorization at:
// https://github.com/mogita/FanFouAPIDoc/wiki/Oauth#%E4%BD%BF%E7%94%A8pin%E7%A0%81%E8%8E%B7%E5%BE%97%E6%8E%88%E6%9D%83
func (c *Client) GetRequestTokenAndURL(callbackURL string) (*RequestToken, string, error) {
	rToken, loginURL, err := c.oauthConsumer.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		return nil, "", CheckAuthResponse(err, "GetRequestTokenAndURL")
	}

	if callbackURL == "oob" {
		loginURL += "&oauth_callback=oob"
	}

	newRequestToken := RequestToken{
		Token:  rToken.Token,
		Secret: rToken.Secret,
	}

	return &newRequestToken, loginURL, nil
}

// AuthorizeClient completes the OAuth authorization to the client
// so it can communicate with Fanfou API
//
// If you use "oob" mode, you also need to provide the verificationCode
func (c *Client) AuthorizeClient(requestToken *RequestToken, verificationCode string) (*AccessToken, error) {
	rToken := oauth.RequestToken{
		Token:  requestToken.Token,
		Secret: requestToken.Secret,
	}

	accessToken, err := c.oauthConsumer.AuthorizeToken(&rToken, verificationCode)

	if err != nil {
		return nil, CheckAuthResponse(err, "AuthorizeClient")
	}

	c.client, err = c.oauthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return nil, err
	}

	newAccessToken := AccessToken{
		Token:          accessToken.Token,
		Secret:         accessToken.Secret,
		AdditionalData: accessToken.AdditionalData,
	}

	return &newAccessToken, nil
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
		return CheckAuthResponse(err, "AuthorizeClientWithXAuth")
	}

	c.client, err = c.oauthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}

// NewRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
//
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

// NewUploadRequest creates an API request dedicated to image uploads.
//
// A relative URL can be provided in uri, in which case it is resolved
// relative to the BaseURL of the Client.
//
// Relative URLs should always be specified without a preceding slash.
func (c *Client) NewUploadRequest(method, uri string, params map[string]string, fileParamName, filePath string) (*http.Request, error) {
	rel, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileParamName, fi.Name())
	if err != nil {
		return nil, err
	}

	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
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
		return nil, err
	}

	response := new(Response)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	tempStr := string(bodyBytes)
	response.BodyStrPtr = &tempStr

	if v != nil {
		response.Data = v
		err = json.Unmarshal(bodyBytes, response.Data)
		c.Response = response
	}

	return response, err
}

// Response specifies Fanfou's response structure.
type Response struct {
	Response   *http.Response // HTTP response
	BodyStrPtr *string
	Data       interface{}
	Meta       *ResponseMeta
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
		Error:   "api request error",
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

// CheckAuthResponse checks the API response for error for
// the requests during authorization, and returns it if present.
//
// A response is considered an error if it has non StatusOK code.
func CheckAuthResponse(err error, tag string) error {
	r := new(ErrorResponse)
	r.Response = &http.Response{StatusCode: http.StatusBadRequest}
	r.Meta = &ResponseMeta{
		Error:   "authorization error",
		Request: "",
	}

	if tag != "" {
		r.Meta.Error = tag + " error: " + err.Error()
	}

	if err, ok := err.(oauth.HTTPExecuteError); ok {
		r.Response = err.Response

		if err.Response.StatusCode >= http.StatusInternalServerError {
			r.Meta.Error = http.StatusText(err.Response.StatusCode)
			return r
		}

		// Fanfou auth errors with a valid body shall be in XML
		decoder := xml.NewDecoder(strings.NewReader(string(err.ResponseBodyBytes)))
		decoder.Strict = false

		if err := decoder.Decode(&r.Meta); err != nil {
			r.Meta.Error = err.Error()
		}
	}

	return r
}
