package fanfou

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	requestTokenURL   = "http://fanfou.com/oauth/request_token"
	authorizeTokenURL = "http://fanfou.com/oauth/authorize"
	accessTokenURL    = "http://fanfou.com/oauth/access_token"
)

const (
	apiBase     = "http://api.fanfou.com"
	apiUserShow = apiBase + "/users/show.json"
)

type baseClient struct {
	HTTPConn *http.Client
}

func (client *baseClient) BaseQuery(queryString string) ([]byte, error) {
	if client.HTTPConn == nil {
		return nil, errors.New("No Client OAuth")
	}

	response, err := client.HTTPConn.Get(queryString)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	return bits, err
}

func (client *baseClient) UserShow(id string) (responseUser, []byte, error) {
	requestURL := fmt.Sprintf("%s?id=%s", apiUserShow, id)
	data, err := client.BaseQuery(requestURL)
	ret := responseUser{}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
