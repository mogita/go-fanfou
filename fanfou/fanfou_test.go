package fanfou

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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
		_, err := fmt.Fprint(w, ``)
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

	// Fanfou client configured to use test server
	client = NewClient("test", "test")
	client.BaseURL, _ = url.Parse(server.URL)
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
