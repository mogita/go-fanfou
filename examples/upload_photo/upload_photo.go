package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mogita/go-fanfou/fanfou"
)

func Usage() {
	fmt.Println("This example will post a new status with a photo to the provided user")
	fmt.Println("---")
	fmt.Println("Usage:")
	fmt.Print("go run examples/upload_photo/upload_photo.go")
	fmt.Print("  --consumerkey <consumerKey>")
	fmt.Print("  --consumersecret <consumerSecret>")
	fmt.Print("  --username <username>")
	fmt.Println("  --password <password>")
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

	username := flag.String(
		"username",
		"",
		"Username from Fanfou.")

	password := flag.String(
		"password",
		"",
		"Password from Fanfou.")

	flag.Parse()

	if len(*consumerKey) == 0 || len(*consumerSecret) == 0 || len(*username) == 0 || len(*password) == 0 {
		fmt.Println("Need to set all flags to run this example: consumerKey, consumerSecret, username, password")
		fmt.Println("---")
		Usage()
		os.Exit(1)
	}

	// Step 1: initialize a new client
	c := fanfou.NewClient(*consumerKey, *consumerSecret)

	// Step 2: authorize the client
	err := c.AuthorizeClientWithXAuth(*username, *password)
	if err != nil {
		// All go-fanfou errors are general errors of ErrorResponse type
		// You can either handle them as normal errors
		// or assert the type and get precise fields like below
		if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
			fmt.Printf("authorize client error: %+v", fanfouErr.GetFanfouError())
			return
		}

		fmt.Println(err)
		return
	}

	// Step 3: call the endpoints
	resp, err := c.Photos.Upload("examples/upload_photo/fanfou.jpg", &fanfou.PhotosOptParams{
		Status: "go-fanfou library test",
	})
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
