package fanfou

import (
	"fmt"
	"log"

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

// DoAuth completes the oauth authorization process
func (client *OAuthClient) DoAuth() error {
	requestToken, loginURL, err := client.OAuthConsumer.GetRequestTokenAndUrl("oob")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + loginURL)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here:")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err := client.OAuthConsumer.AuthorizeToken(requestToken, verificationCode)

	if err != nil {
		log.Fatal(err)
	}

	client.http, err = client.OAuthConsumer.MakeHttpClient(accessToken)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
