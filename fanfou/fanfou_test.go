package fanfou

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/mogita/oauth"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Fanfou client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {

	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	mux.HandleFunc("/"+requestTokenURI, func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, `oauth_token=test_token&oauth_token_secret=test_secret`)
		if err != nil {
			panic(fmt.Sprintf("request token URI mock server error: %+v", err))
		}
	})

	mux.HandleFunc("/"+authorizeTokenURI, func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, ``)
		if err != nil {
			panic(fmt.Sprintf("authorize token URI mock server error: %+v", err))
		}
	})

	mux.HandleFunc("/"+accessTokenURI, func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, `oauth_token=test&oauth_token_secret=test`)
		if err != nil {
			panic(fmt.Sprintf("access token URI mock server error: %+v", err))
		}
	})

	// mock base url
	mockBaseURL, _ := url.Parse(server.URL)
	AuthBaseURL = mockBaseURL.String() + "/"

	// Fanfou client configured to use test server
	client = NewClient("test", "test")
	client.BaseURL = mockBaseURL
	err := client.AuthorizeClientWithXAuth("", "")
	if err != nil {
		panic(err)
	}
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	for key, want := range values {
		if v := r.FormValue(key); v != want {
			t.Errorf("Request parameter %v = %v, want %v", key, v, want)
		}
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient("", "")

	want := "http://api.fanfou.com/"
	if c.BaseURL.String() != want {
		t.Errorf("NewClient BaseURL = %v, want %v", c.BaseURL.String(), want)
	}
	want = "github.com/mogita/go-fanfou v0.1"
	if c.UserAgent != want {
		t.Errorf("NewClient UserAgent = %v, want %v", c.UserAgent, want)
	}
}

func TestGetRequestTokenAndURL(t *testing.T) {
	setup()
	defer teardown()

	c := NewClient("", "")
	rToken, loginURL, err := c.GetRequestTokenAndURL("")
	if err != nil {
		panic(err)
	}

	want := reflect.TypeOf(oauth.RequestToken{Token: "", Secret: ""})
	actual := reflect.TypeOf(*rToken)
	if actual != want {
		t.Errorf("NewClient rToken type = %v, want %v", actual, want)
	}

	want = reflect.TypeOf("")
	actual = reflect.TypeOf(loginURL)
	if actual != want {
		t.Errorf("NewClient loginURL type = %v, want %v", actual, want)
	}
}

func TestAuthorizeClient(t *testing.T) {
	setup()
	defer teardown()

	c := NewClient("", "")
	rToken, _, err := c.GetRequestTokenAndURL("")
	if err != nil {
		panic(err)
	}

	err = c.AuthorizeClient(rToken)
	if err != nil {
		panic(err)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient("", "")

	inURL, outURL := "foo/bar.json", c.BaseURL.String()+"foo/bar.json"
	req, _ := c.NewRequest("GET", inURL, "")

	// test that relative URL was expanded and access token appears in query string
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that default user-agent is attached to the requet
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, want %v", userAgent, c.UserAgent)
	}
}

func TestNewUploadRequest(t *testing.T) {
	c := NewClient("", "")

	inURL, outURL := "foo/bar.json", c.BaseURL.String()+"foo/bar.json"
	req, _ := c.NewUploadRequest("POST", inURL, map[string]string{"test_key": "test_value"}, "photo", "./fanfou.go")

	// test that relative URL was expanded and access token appears in query string
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, want %v", inURL, req.URL, outURL)
	}

	// test that default user-agent is attached to the requet
	userAgent := req.Header.Get("User-Agent")
	if c.UserAgent != userAgent {
		t.Errorf("NewRequest() User-Agent = %v, want %v", userAgent, c.UserAgent)
	}
}

func TestCheckResponse(t *testing.T) {
	mockRes := http.Response{
		StatusCode: http.StatusOK,
	}
	err := CheckResponse(&mockRes)
	if err != nil {
		t.Errorf("CheckResponse() while 200, result %v, want %v", err, nil)
	}

	mockRes = http.Response{
		StatusCode: http.StatusBadRequest,
		Request: &http.Request{
			Method: http.MethodPost,
			URL: &url.URL{
				Scheme: "https",
				Host:   "test.url.com",
				Path:   "/",
			},
		},
		Body: ioutil.NopCloser(bytes.NewBufferString(`{"error":"test_error", "request": "test_request"}`)),
	}

	err = CheckResponse(&mockRes)
	if err == nil {
		t.Errorf("CheckResponse() while 400, result %v, want err", err)
	}

	want := ""
	actual := err.Error()
	if reflect.TypeOf(actual) != reflect.TypeOf(want) {
		t.Errorf("CheckResponse() while 400, err.Error() type is %v, want %v", reflect.TypeOf(actual), reflect.TypeOf(want))
	}

	fanfouErr, ok := err.(*ErrorResponse)
	if !ok {
		t.Errorf("CheckResponse() while 400, error is not ErrorResponse, want ErrorResponse")
	}

	want = "400"
	actual = fanfouErr.GetStatusCode()
	if want != actual {
		t.Errorf("CheckResponse() while 400, fanfouErr.GetStatusCode() is %v, want %v", actual, want)
	}

	want = "POST"
	actual = fanfouErr.GetRequestMethod()
	if want != actual {
		t.Errorf("CheckResponse() while 400, fanfouErr.GetRequestMethod() is %v, want %v", actual, want)
	}

	want = "https://test.url.com/"
	actual = fanfouErr.GetRequestURL()
	if want != actual {
		t.Errorf("CheckResponse() while 400, fanfouErr.GetRequestURL() is %v, want %v", actual, want)
	}

	want = "test_error"
	actual = fanfouErr.GetFanfouError()
	if want != actual {
		t.Errorf("CheckResponse() while 400, fanfouErr.GetFanfouError() is %v, want %v", actual, want)
	}
}
