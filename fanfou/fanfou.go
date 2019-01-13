package fanfou

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mrjones/oauth"
)

const (
	// LibraryVersion represents this library version
	LibraryVersion = "0.1"

	// UserAgent represents this client User-Agent
	UserAgent = "github.com/mogita/go-fanfou v" + LibraryVersion

	// BaseURL represents Fanfou API base URL
	BaseURL = "http://api.fanfou.com/"

	// Request token URL
	requestTokenURL = "http://fanfou.com/oauth/request_token"

	// Authorize token URL
	authorizeTokenURL = "http://fanfou.com/oauth/authorize"

	// Access token URL
	accessTokenURL = "http://fanfou.com/oauth/access_token"
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
	// Users         *UsersService
	// Relationships *RelationshipsService
	// Media         *MediaService
	// Comments      *CommentsService
	// Likes         *LikesService
	// Tags          *TagsService
	// Locations     *LocationsService
	// Geographies   *GeographiesService
	Trends *TrendsService

	Temporary Response
	Response  *Response
}

// Response specifies Fanfou's response structure.
type Response struct {
	Response     *http.Response // HTTP response
	Data         interface{}    `json:"data,omitempty"`          // business data
	Meta         interface{}    `json:"meta,omitempty"`          // meta info
	ErrorMessage *ErrorMessage  `json:"error_message,omitempty"` // carries error information if any
}

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Request string `json:"request,omitempty"`
	Error   string `json:"error,omitempty"`
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
				RequestTokenUrl:   requestTokenURL,
				AuthorizeTokenUrl: authorizeTokenURL,
				AccessTokenUrl:    accessTokenURL,
			},
		),
	}

	c.oauthConsumer.Debug(false)

	// c.Users = &UsersService{client: c}
	// c.Relationships = &RelationshipsService{client: c}
	// c.Media = &MediaService{client: c}
	// c.Comments = &CommentsService{client: c}
	// c.Likes = &LikesService{client: c}
	// c.Tags = &TagsService{client: c}
	// c.Locations = &LocationsService{client: c}
	// c.Geographies = &GeographiesService{client: c}
	c.Trends = &TrendsService{client: c}

	return c
}

// GetRequestTokenAndUrl returns the request token and the login url for authorizing this token
func (c *Client) GetRequestTokenAndUrl(callbackURL string) (*oauth.RequestToken, string, error) {
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
		return err
	}

	c.client, err = c.oauthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AuthorizeClientWithXAuth(username, password string) error {
	c.oauthConsumer.AdditionalParams["x_auth_username"] = username
	c.oauthConsumer.AdditionalParams["x_auth_password"] = password
	c.oauthConsumer.AdditionalParams["x_auth_mode"] = "client_auth"

	reqToken := oauth.RequestToken{}
	accessToken, err := c.oauthConsumer.AuthorizeToken(&reqToken, "")

	if err != nil {
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

	r := &Response{Response: resp}
	if v != nil {
		r.Data = v
		err = json.NewDecoder(resp.Body).Decode(v)
		c.Response = r
	}

	return resp, err
}

// ErrorResponse represents a Response which contains an error
type ErrorResponse Response

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorMessage.Error)
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(res *http.Response) error {
	if res.StatusCode == http.StatusOK {
		return nil
	}

	resp := new(ErrorResponse)
	resp.Response = res

	if res.StatusCode == http.StatusInternalServerError || res.StatusCode == http.StatusNotFound {
		return resp
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		resp.ErrorMessage.Error = err.Error()
	}
	if err := json.Unmarshal(data, resp); err != nil {
		resp.ErrorMessage.Error = err.Error()
	}

	return resp
}
