<p align="center">
  <img src="fanfou.png?raw=true" width="300" height="200" />
  <h3 align="center">GO FANFOU</h3>
  <p align="center">A Fanfou API library for <a href="http://golang.org/" target="_blank">Go</a></p>
  <p align="center">
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/mogita/go-fanfou"><img alt="Travis" src="https://img.shields.io/travis/mogita/go-fanfou/master.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/mogita/go-fanfou"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/mogita/go-fanfou?style=flat-square"></a>
    <a href="https://coveralls.io/github/mogita/go-fanfou?branch=master"><img alt="Coverage Status" src="https://img.shields.io/coveralls/mogita/go-fanfou/master.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/mogita/go-fanfou/fanfou"><img alt="GoDoc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
</p>

## Key Features

- Simple and intuitive endponits and error wrapping
- Struct and original JSON output
- Supports OAuth (safer) and XAuth (simpler)
- Covers all endpoints of [Fanfou API v1](https://github.com/mogita/FanFouAPIDoc/wiki)

## Installation

```
$ go get -u github.com/mogita/go-fanfou
```

## Usage

```go
package main

import "github.com/mogita/go-fanfou/fanfou"

func main() {
  // ...
}
```

### Basic

To call an endpoint e.g. `/statuses/public_timeline`, you can call it like this:

```go
// Every API endpoint has the same return value structure
data, JSON, err := c.Statuses.PublicTimeline(&fanfou.StatusesOptParams{
    Count: 10,
})
```

All optional parameter types starts with the resource's name. E.g. `Statuses` -> `StatusesOptParams`.

See the `examples` directory to learn how to authenticate the client instance before calling the endpoints.

### Error Handling

Errors default to the format as below:

```
POST http://api.fanfou.com/photos/upload.json: 400 上传照片失败
```
 
Meanwhile they can be asserted to extract the specific detail that you can use to handle the errors programmatically. Like this:

```go
_, _, err := c.Statuses.PublicTimeline(nil)

if err != nil {
    if fanfouErr, ok := err.(*fanfou.ErrorResponse); ok {
    	// Will print only the error message text returned by Fanfou API
        fmt.Printf("%s\n", fanfouErr.GetFanfouError())
        return
    }

    // Will print the default error format
    fmt.Println(err)
    return
}
```

## Running the Examples

Check out the `examples` folder for working code snippets. You can run the examples with these commands to see how this library works:

### Standard OAuth

```shell
$ go run examples/oauth/oauth.go --consumerkey <your_consumer_key> --consumersecret <your_consumer_secret>
```

### OOB OAuth

> OOB stands for "out of band"

```shell
$ go run examples/oauth_oob/oauth_oob.go --consumerkey <your_consumer_key> --consumersecret <your_consumer_secret>
```

### XAuth

```shell
$ go run examples/xauth/xauth.go --consumerkey <your_consumer_key> --consumersecret <your_consumer_secret> --username <your_username> --password <your_password>
```

### Upload Photos

```shell
$ go run examples/upload_photo/upload_photo.go --consumerkey <your_consumer_key> --consumersecret <your_consumer_secret> --username <your_username> --password <your_password>
```

## Credits

- [oauth](https://godoc.org/github.com/mogita/oauth) (a fork of [mrjones/oauth](https://godoc.org/github.com/mrjones/oauth)) - OAuth 1.0 implementation in go (golang)
- [go-github](https://github.com/google/go-github) - This library mimics its structure. A copy of its LICENSE can be found here [go-github-LICENSE](./go-github-LICENSE)

## Contributing

Thank you very much for paying attention to this library. If you feel like helping improve it, please kindly make sure to follow the instructions:

Link the pre-commit hook which runs tests and go-fmt before committing

```
ln -s $PWD/pre-commit.sh .git/hooks/pre-commit
```

Always run tests before committing

```
go test ./...
```

You can also follow this [Trello board](https://trello.com/b/Z6XTVn7U/go-fanfou) if you're interested in the progress of this project and its sibling products.

## License

MIT © [mogita](https://github.com/mogita)
