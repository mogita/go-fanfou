package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/mogita/go-fanfou/fanfou"
)

var tokens map[string]*fanfou.RequestToken
var c *fanfou.Client

func main() {
	tokens = make(map[string]*fanfou.RequestToken)

	// Starting the server with a client instance
	consumerKey := flag.String(
		"consumerkey",
		"",
		"Consumer Key from Fanfou. See: https://fanfou.com/apps")

	consumerSecret := flag.String(
		"consumersecret",
		"",
		"Consumer Secret from Fanfou. See: https://fanfou.com/apps")

	port := flag.Int(
		"port",
		8080,
		"Port to listen on.")

	flag.Parse()

	c = fanfou.NewClient(*consumerKey, *consumerSecret)

	http.HandleFunc("/", RedirectUserToFanfou)
	http.HandleFunc("/callback", GetFanfouToken)
	u := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening on '%s'\n", u)
	fmt.Println("You can visit http://localhost:8080 to start the authorization process")
	err := http.ListenAndServe(u, nil)
	if err != nil {
		panic(err)
	}
}

func RedirectUserToFanfou(w http.ResponseWriter, r *http.Request) {
	tokenUrl := fmt.Sprintf("http://%s/callback", r.Host)
	token, requestUrl, err := c.GetRequestTokenAndURL(tokenUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to save the token, we'll need it for AuthorizeToken()
	tokens[token.Token] = &fanfou.RequestToken{
		Token:  token.Token,
		Secret: token.Secret,
	}

	http.Redirect(w, r, requestUrl, http.StatusTemporaryRedirect)
}

func GetFanfouToken(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	accessToken, err := c.AuthorizeClient(tokens[tokenKey], verificationCode)
	if err != nil {
		panic(err)
	}

	// You can distinguish your users with the information in the accessToken
	fmt.Println(accessToken)

	// Calling the endpoints
	resp, err := c.Statuses.HomeTimeline(&fanfou.StatusesOptParams{
		Count:  3,
		Format: "html",
	})

	if err != nil {
		if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
			fmt.Printf("%s\n", fanfouErr.GetFanfouError())
			return
		}

		fmt.Println(err)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Fprintf(w, string(respJSON))
	if err != nil {
		panic(err)
	}
}
