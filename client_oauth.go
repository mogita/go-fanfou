package fanfou

import (
	"github.com/mrjones/oauth"
)

// OAuthClient wraps a base client and an oauth consumer
type OAuthClient struct {
	baseClient
	OAuthConsumer *oauth.Consumer
}

// NewClientWithOAuth returns a client without authorization
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

// DoAuthOob completes the oauth authorization process
func (client *OAuthClient) DoAuthOob(rToken *oauth.RequestToken, verificationCode string) error {
	accessToken, err := client.OAuthConsumer.AuthorizeToken(rToken, verificationCode)

	if err != nil {
		return err
	}

	client.http, err = client.OAuthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		return err
	}

	return nil
}
