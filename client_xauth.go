package fanfou

import (
	"log"

	"github.com/mrjones/oauth"
)

// XAuthClient wraps a base client and an oauth consumer
type XAuthClient struct {
	baseClient
	oAuthConsumer *oauth.Consumer
}

// NewClientWithXAuth returns an authorized client
func NewClientWithXAuth(consumerKey, consumerSecret, username, password string) (*XAuthClient, error) {
	newClient := new(XAuthClient)
	newClient.oAuthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authorizeTokenURL,
			AccessTokenUrl:    accessTokenURL,
		},
	)

	newClient.oAuthConsumer.AdditionalParams["x_auth_username"] = username
	newClient.oAuthConsumer.AdditionalParams["x_auth_password"] = password
	newClient.oAuthConsumer.AdditionalParams["x_auth_mode"] = "client_auth"
	newClient.oAuthConsumer.Debug(false)

	err := newClient.doXAuth()

	if err != nil {
		return nil, err
	}

	return newClient, nil
}

func (client *XAuthClient) doXAuth() error {
	reqToken := oauth.RequestToken{}
	accessToken, err := client.oAuthConsumer.AuthorizeToken(&reqToken, "")

	if err != nil {
		log.Fatal(err)
	}

	client.http, err = client.oAuthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
