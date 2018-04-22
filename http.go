package fanfou

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

const (
	requestTokenURL   = "http://fanfou.com/oauth/request_token"
	authorizeTokenURL = "http://fanfou.com/oauth/authorize"
	accessTokenURL    = "http://fanfou.com/oauth/access_token"
)

type baseClient struct {
	http *http.Client
}

func (client *baseClient) makeRequest(method, path string, params *ReqParams) (byteData []byte, err error) {
	var resp *http.Response

	if client.http == nil {
		return nil, errors.New("Invalid OAuth client")
	}

	switch method {
	case http.MethodGet:
		paramValues := client.paramsToURLValues(params)
		queryString := paramValues.Encode()
		requestPath := fmt.Sprintf("%s?%s", path, queryString)
		resp, err = client.http.Get(requestPath)
	case http.MethodPost:
		resp, err = client.http.PostForm(path, client.paramsToURLValues(params))
	case "photo":
		// invoked by photos upload
		req, nfurRrr := client.newfileUploadRequest(apiAccountUpdateProfileImage, map[string]string{"status": params.Status}, "photo", params.Photo)
		if nfurRrr != nil {
			return nil, fmt.Errorf("Could not initialize the photos upload request: %#v", nfurRrr)
		}

		resp, err = client.http.Do(req)
	case "image":
		// invoked by account update profile image
		req, nfurRrr := client.newfileUploadRequest(apiAccountUpdateProfileImage, map[string]string{"status": params.Status}, "image", params.Image)
		if nfurRrr != nil {
			return nil, fmt.Errorf("Could not initialize the image upload request: %#v", nfurRrr)
		}

		resp, err = client.http.Do(req)
	default:
		return nil, fmt.Errorf("Unsupported request method: %#v", method)
	}

	if err != nil {
		return nil, fmt.Errorf("Could not make request: %#v", err)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		bits, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("Could not read response body: %#v", err)
		}

		return []byte(strings.TrimSpace(string(bits))), nil
	case http.StatusAccepted:
		// according to what the docs describes, 202 should be almost obsolete
		// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/ErrorCodes_Responses#http-status-codes
		return nil, nil
	case http.StatusBadRequest:
		return nil, errors.New("Bad Request")
	case http.StatusUnauthorized:
		return nil, errors.New("Unauthorized")
	case http.StatusForbidden:
		return nil, errors.New("Forbidden")
	case http.StatusNotFound:
		return nil, errors.New("Not Found")
	default:
		return nil, errors.New("General Error Response")
	}
}

func (client *baseClient) paramsToJSON(params *ReqParams) (*string, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("Could not marshal json: %#v", err)
	}
	jsonStr := string(json)
	return &jsonStr, nil
}

func (client *baseClient) paramsToURLValues(params *ReqParams) (values url.Values) {
	values = url.Values{}

	paramsVals := reflect.ValueOf(params).Elem()
	typeVals := paramsVals.Type()

	for i := 0; i < paramsVals.NumField(); i++ {
		field := paramsVals.Field(i)
		tag := typeVals.Field(i).Tag.Get("json")
		value := field.String()

		if value != "" {
			values.Set(strings.Replace(tag, ",omitempty", "", -1), value)
		}
	}

	return
}

func (client *baseClient) newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
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

	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}

	part.Write(fileContents)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}
