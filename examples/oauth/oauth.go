package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mrjones/oauth"

	"github.com/mogita/go-fanfou/examples"
	"github.com/mogita/go-fanfou/fanfou"
)

// modify the credentials in def.go to your own keys etc.
const (
	consumerKey    = examples.ConsumerKey
	consumerSecret = examples.ConsumerSecret
	cbURL          = examples.CallbackURL
)

func main() {
	tokens := map[string]*oauth.RequestToken{}
	client := fanfou.NewClientWithOAuth(consumerKey, consumerSecret)

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		rToken, authURL, err := client.GetTokenAndAuthURL(cbURL)

		if err != nil {
			log.Println(fmt.Errorf("could not get request token and auth url: %+v", err))
		}

		// the map is a way to store sessions
		// you may want to implement your app in a more secure way
		tokens[rToken.Token] = rToken

		html := fmt.Sprintf("<a href=\"%s\">%s</a>", authURL, "Click to authorize")
		_, err = resp.Write([]byte(html))
		if err != nil {
			fmt.Printf("%+v", err)
		}
	})

	http.HandleFunc("/callback", func(resp http.ResponseWriter, req *http.Request) {
		values := req.URL.Query()
		token := values.Get("oauth_token")

		err := client.DoAuth(tokens[token])
		if err != nil {
			log.Println(fmt.Errorf("could not do auth: %+v", err))
		}

		user, _, err := client.AccountVerifyCredentials(&fanfou.ReqParams{})
		if err != nil {
			log.Println(fmt.Errorf("could not make request: %+v", err))
		}

		jsonBytes, err := json.Marshal(user)
		if err != nil {
			log.Println(fmt.Errorf("could not marshal json: %+v", err))
		}

		_, err = resp.Write(jsonBytes)
		if err != nil {
			fmt.Printf("%+v", err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
