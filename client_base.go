package fanfou

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	requestTokenURL   = "http://fanfou.com/oauth/request_token"
	authorizeTokenURL = "http://fanfou.com/oauth/authorize"
	accessTokenURL    = "http://fanfou.com/oauth/access_token"
)

const (
	apiBase           = "http://api.fanfou.com"
	apiUserShow       = apiBase + "/users/show.json"
	apiStatusesUpdate = apiBase + "/statuses/update.json"
)

type baseClient struct {
	HTTPConn *http.Client
}

func (client *baseClient) query(method, path string, params map[string]string) (byteData []byte, err error) {
	var resp *http.Response

	if client.HTTPConn == nil {
		return nil, errors.New("No Client OAuth")
	}

	switch method {
	case http.MethodGet:
		resp, err = client.HTTPConn.Get(path)
	case http.MethodPost:
		resp, err = client.HTTPConn.PostForm(path, client.paramsToURLValues(params))
	default:
		return nil, fmt.Errorf("Unsupported http method: %#v", method)
	}

	if err != nil {
		return nil, fmt.Errorf("Could not make http request: %#v", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		bits, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("Could not read response body: %#v", err)
		}

		return bits, nil
	case http.StatusBadRequest:
		return nil, errors.New("Bad Request")
	case http.StatusNotFound:
		return nil, errors.New("Not Found")
	default:
		return nil, errors.New("Other error")
	}
}

func (client *baseClient) paramsToURLValues(params map[string]string) url.Values {
	output := url.Values{}

	for key, value := range params {
		output[key] = []string{value}
	}

	return output
}

func (client *baseClient) UserShow(id string) (*responseUser, []byte, error) {
	requestURL := fmt.Sprintf("%s?id=%s", apiUserShow, id)
	data, err := client.query(http.MethodGet, requestURL, nil)

	if err != nil {
		return nil, nil, err
	}

	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return &ret, data, err
}

func (client *baseClient) StatusesUpdate(status string, inReplyToStatusID, inReplyToUserID, repostStatusID, location *string) (*responseStatus, []byte, error) {
	params := map[string]string{
		"status": status,
	}

	if inReplyToStatusID != nil {
		params["in_reply_to_status_id"] = *inReplyToStatusID
	}

	if inReplyToUserID != nil {
		params["in_reply_to_user_id"] = *inReplyToUserID
	}

	if repostStatusID != nil {
		params["repost_status_id"] = *repostStatusID
	}

	if location != nil {
		params["location"] = *location
	}

	data, err := client.query(http.MethodPost, apiStatusesUpdate, params)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to exec query: %+v", err)
	}

	ret := responseStatus{}
	dataTrim := strings.TrimSpace(string(data))
	err = json.Unmarshal([]byte(dataTrim), &ret)
	return &ret, data, err
}
