package main

import (
	"fmt"

	"github.com/mogita/go-fanfou/examples"
	"github.com/mogita/go-fanfou/fanfou"
)

const (
	consumerKey    = examples.ConsumerKey
	consumerSecret = examples.ConsumerSecret
	username       = examples.Username
	password       = examples.Password
)

func main() {
	c := fanfou.NewClient(consumerKey, consumerSecret)

	err := c.AuthorizeClientWithXAuth(username, password)
	if err != nil {
		if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
			fmt.Printf("authorize client error: %+v", fanfouErr.GetFanfouError())
			return
		}

		fmt.Println(err)
		return
	}

	resp, err := c.Statuses.HomeTimeline(&fanfou.StatusesOptParams{ID: "montanathief", Count: 2, Format: "html"})
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
