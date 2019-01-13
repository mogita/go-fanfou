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
		panic(fmt.Sprintf("authorize client error: %+v", err))
	}

	trends, err := c.Trends.List()
	if err != nil {
		fmt.Printf("as_of: %+v\n", err)
	}

	fmt.Printf("trends as_of: %s\n", trends.AsOf)

	for index, trend := range trends.Trends {
		fmt.Printf("trend %d query: %s\n", index, trend.Query)
		fmt.Printf("trend %d name: %s\n", index, trend.Name)
		fmt.Printf("trend %d url: %s\n", index, trend.URL)
	}
}
