// Package fanfou is an API client SDK for fanfou.com
package fanfou

import (
	"github.com/mrjones/oauth"
)

// OAuthClient is the core type
type OAuthClient struct {
	httpClientWrapper
	OAuthConsumer *oauth.Consumer
}

// NewClientWithOAuth returns a client without authorization. In order to authorize the client, you need to:
//
// 1) Call "GetTokenAndAuthURL" method to obtain a request token and the auth URL. Let your user access the auth URL and grant access to the request token.
//
// 2) Call "DoAuth" method to complete the auth process after the request token has been authorized by the user.
func NewClientWithOAuth(consumerKey, consumerSecret string) *OAuthClient {
	newClient := new(OAuthClient)
	newClient.OAuthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authorizeTokenURL,
			AccessTokenUrl:    accessTokenURL,
		},
	)

	newClient.OAuthConsumer.Debug(false)

	return newClient
}

// NewClientWithXAuth returns an authorized client. The client is ready to call all API methods immediately.
//
// For more information on XAuth authorization method, please refer to https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Xauth
func NewClientWithXAuth(consumerKey, consumerSecret, username, password string) (*OAuthClient, error) {
	newClient := NewClientWithOAuth(consumerKey, consumerSecret)

	newClient.OAuthConsumer.AdditionalParams["x_auth_username"] = username
	newClient.OAuthConsumer.AdditionalParams["x_auth_password"] = password
	newClient.OAuthConsumer.AdditionalParams["x_auth_mode"] = "client_auth"

	err := newClient.doXAuth()

	if err != nil {
		return nil, err
	}

	return newClient, nil
}

// GetTokenAndAuthURL returns the request token and the url for authorizing this token
func (client *OAuthClient) GetTokenAndAuthURL(callbackURL string) (*oauth.RequestToken, string, error) {
	rToken, authURL, err := client.OAuthConsumer.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		return nil, "", err
	}

	return rToken, authURL, nil
}

// DoAuth completes the oauth authorization process
func (client *OAuthClient) DoAuth(rToken *oauth.RequestToken) error {
	accessToken, err := client.OAuthConsumer.AuthorizeToken(rToken, "")

	if err != nil {
		return err
	}

	client.http, err = client.OAuthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}

func (client *OAuthClient) doXAuth() error {
	reqToken := oauth.RequestToken{}
	accessToken, err := client.OAuthConsumer.AuthorizeToken(&reqToken, "")

	if err != nil {
		return err
	}

	client.http, err = client.OAuthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}
