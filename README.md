# FANFOU GO

A Fanfou API Client for Gophers

## Usage

```go
package main

import (
  "git.mogita.com/mogita/go-fanfou"
)

const (
	consumerKey    = "xxx"
	consumerSecret = "xxx"
	username       = "xxx"
	password       = "xxx"
)

func main() {
	client, _ := fanfou.NewClientWithXAuth(consumerKey, consumerSecret, username, password)

	res, _, err := client.UserShow(&fanfou.ReqParams{
		ID: "mogita",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n\n", res)

	resS, _, err := client.StatusesUpdate(&fanfou.ReqParams{
		Status: "niubi",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resS)
}
```

## License

MIT © [mogita](https://github.com/mogita)
