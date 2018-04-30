package main

import (
	"fmt"
	"log"

	"git.mogita.com/mogita/go-fanfou/examples"
	"git.mogita.com/mogita/go-fanfou/fanfou"
)

// modify the credentials in def.go to your own keys etc.
const (
	consumerKey    = examples.ConsumerKey
	consumerSecret = examples.ConsumerSecret
	username       = examples.Username
	password       = examples.Password
)

func main() {
	client, _ := fanfou.NewClientWithXAuth(consumerKey, consumerSecret, username, password)

	res, _, err := client.PhotosUpload(&fanfou.ReqParams{Status: "I'm uploading a photo", Photo: "./sample.jpg"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", res)
}
