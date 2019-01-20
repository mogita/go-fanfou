<p align="center">
  <img src="fanfou.png?raw=true" width="300" height="200" />
  <h2 align="center">GO FANFOU</h2>
  <p align="center">A Fanfou API library for <a href="http://golang.org/" target="_blank">Go</a></p>
  <p align="center">
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/mogita/go-fanfou"><img alt="Travis" src="https://img.shields.io/travis/mogita/go-fanfou/master.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/mogita/go-fanfou"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/mogita/go-fanfou?style=flat-square"></a>
    <a href="https://coveralls.io/github/mogita/go-fanfou?branch=master"><img alt="Coverage Status" src="https://img.shields.io/coveralls/mogita/go-fanfou/master.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/mogita/go-fanfou/fanfou"><img alt="GoDoc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
</p>

---

## Key Features

- Simple and intuitive endponits and error wrapping
- Struct and original JSON output
- Supports OAuth (safer) and XAuth (simpler)
- Covers all endpoints of [Fanfou API v1](https://github.com/mogita/FanFouAPIDoc/wiki)

## Installation

```
go get -u github.com/mogita/go-fanfou
```

## Usage

```go
package main

import "github.com/mogita/go-fanfou/fanfou"

func main() {
  // ...
}
```

Please refer to the `examples` folder for the basic usages of this library. You can run the examples with these commands to see how this library works:

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

## Built With

- [oauth](https://godoc.org/github.com/mogita/oauth) (a fork of [mrjones/oauth](https://godoc.org/github.com/mrjones/oauth)) - OAuth 1.0 implementation in go (golang)

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

You can also follow this [Trello board](https://trello.com/b/Z6XTVn7U/go-fanfou) if you're interested in the progress of this project and also its sibling products.

## License

MIT © [mogita](https://github.com/mogita)
