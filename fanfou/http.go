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

type httpClientWrapper struct {
	http *http.Client
}

func (client *httpClientWrapper) makeRequest(method, path string, params *ReqParams) (byteData []byte, err error) {
	var resp *http.Response

	if client.http == nil {
		return nil, errors.New("Invalid OAuth client")
	}

	switch method {
	case http.MethodGet:
		paramValues := paramsToURLValues(params)
		queryString := paramValues.Encode()
		requestPath := fmt.Sprintf("%s?%s", path, queryString)
		resp, err = client.http.Get(requestPath)
	case http.MethodPost:
		resp, err = client.http.PostForm(path, paramsToURLValues(params))
	case "photo":
		// invoked by photos upload
		req, nfurRrr := newfileUploadRequest(endpoints["PhotosUpload"], map[string]string{"status": params.Status}, "photo", params.Photo)
		if nfurRrr != nil {
			return nil, fmt.Errorf("Could not initialize the photos upload request: %#v", nfurRrr)
		}

		resp, err = client.http.Do(req)
	case "image":
		// invoked by account update profile image
		req, nfurRrr := newfileUploadRequest(endpoints["AccountUpdateProfileImage"], map[string]string{"status": params.Status}, "image", params.Image)
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

	// read response body data
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response body: %#v", err)
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest {
		// return nice responses
		return []byte(strings.TrimSpace(string(respBodyBytes))), nil
	}

	// process and return bad responses
	respErr := responseError{}
	if err = json.Unmarshal(respBodyBytes, &respErr); err != nil {
		return nil, fmt.Errorf("Malformed error response body: %+v. Original error: [%d] %#v", err, resp.StatusCode, string(respBodyBytes))
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, fmt.Errorf("Bad Request: %s. %s", respErr.Request, respErr.Error)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Unauthorized: %s. %s", respErr.Request, respErr.Error)
	case http.StatusForbidden:
		return nil, fmt.Errorf("Forbidden: %s. %s", respErr.Request, respErr.Error)
	case http.StatusNotFound:
		return nil, fmt.Errorf("Not Found: %s. %s", respErr.Request, respErr.Error)
	default:
		return nil, fmt.Errorf("Other errors: %s. %s", respErr.Request, respErr.Error)
	}
}

func paramsToURLValues(params *ReqParams) (values url.Values) {
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

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
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
