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

	res, _, err := client.UserShow("mogita")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", res)
}
```

## License

MIT © [mogita](https://github.com/mogita)
