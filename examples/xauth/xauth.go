package main

import (
	"fmt"
	"log"

	"github.com/mogita/go-fanfou/examples"
	"github.com/mogita/go-fanfou/fanfou"
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

	res, _, err := client.SearchPublicTimeline(&fanfou.ReqParams{Q: "有猫"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", res)
}
