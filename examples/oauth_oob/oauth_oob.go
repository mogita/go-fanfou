package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mogita/go-fanfou/fanfou"
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Print("go run examples/oauth_oob/oauth_oob.go")
	fmt.Print("  --consumerkey <consumerKey>")
	fmt.Println("  --consumersecret <consumerSecret>")
	fmt.Println("")
	fmt.Println("In order to get your consumerKey and consumerSecret, you must register an 'app' at fanfou.com:")
	fmt.Println("https://fanfou.com/apps")
}

func main() {
	// Antecedent steps
	consumerKey := flag.String(
		"consumerkey",
		"",
		"Consumer Key from Fanfou. See: https://fanfou.com/apps")

	consumerSecret := flag.String(
		"consumersecret",
		"",
		"Consumer Secret from Fanfou. See: https://fanfou.com/apps")

	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 {
		fmt.Println("Need to set all flags to run this example: consumerKey, consumerSecret")
		fmt.Println("---")
		Usage()
		os.Exit(1)
	}

	// Step 1: initialize a new client
	c := fanfou.NewClient(*consumerKey, *consumerSecret)

	// Step 2: authorize the client
	requestToken, URL, err := c.GetRequestTokenAndURL("oob")
	if err != nil {
		// All go-fanfou errors are general errors of ErrorResponse type
		// You can either handle them as normal errors
		// or assert the type and get precise fields like below
		if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
			fmt.Printf("authorize client error: %+v", fanfouErr.Error())
			return
		}

		fmt.Println(err)
		return
	}

	fmt.Println("(1) Go to: " + URL)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	_, err = fmt.Scanln(&verificationCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken, err := c.AuthorizeClient(requestToken, verificationCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// You can distinguish your users with the information in the accessToken
	fmt.Println(accessToken)

	// Step 3: call the endpoints
	resp, _, err := c.Statuses.HomeTimeline(&fanfou.StatusesOptParams{Count: 3, Format: "html"})
	if err != nil {
		if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
			fmt.Printf("%s\n", fanfouErr.GetFanfouError())
			return
		}

		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp)
}
