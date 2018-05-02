package fanfou

import (
	"fmt"
	"net/http"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

const (
	mockConsumerKey    = "mock_consumer_key"
	mockConsumerSecret = "mock_consumer_secret"
	mockRequestToken   = "mock_request_token"
	mockRequestSecret  = "mock_request_secret"
	mockAccessToken    = "mock_access_token"
	mockAccessSecret   = "mock_access_secret"
)

func init() {
	fmt.Println("test starts")

	for _, value := range endpoints {
		httpmock.RegisterResponder("GET", value, func(req *http.Request) (*http.Response, error) {
			res, _ := httpmock.NewJsonResponse(200, `{"id":"mogita", "reset_time":"some_time", "remaining_hits": 1, "hourly_limit": 1, "reset_time_in_seconds": 1, "mentions": 1, "direct_messages": 1, "friend_requests": 1, "result": "test_result", "notify_num": "test_notify_num", "name": "test_name", "as_of": "test_as_of", "relationship": {}, "msg_num": 1}`)
			return res, nil
		})

		httpmock.RegisterResponder("POST", value, func(req *http.Request) (*http.Response, error) {
			res, _ := httpmock.NewJsonResponse(200, `{"id":"mogita", "reset_time":"some_time", "remaining_hits": 1, "hourly_limit": 1, "reset_time_in_seconds": 1, "mentions": 1, "direct_messages": 1, "friend_requests": 1, "result": "test_result", "notify_num": "test_notify_num", "name": "test_name", "as_of": "test_as_of", "relationship": {}, "msg_num": 1}`)
			return res, nil
		})
	}

	httpmock.RegisterResponder("GET", requestTokenURL, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, fmt.Sprintf(`oauth_token=%s&oauth_token_secret=%s`, mockRequestToken, mockRequestSecret)), nil
	})

	httpmock.RegisterResponder("GET", accessTokenURL, func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(200, fmt.Sprintf(`oauth_token=%s&oauth_token_secret=%s`, mockAccessToken, mockAccessSecret)), nil
	})

	httpmock.Activate()
}
