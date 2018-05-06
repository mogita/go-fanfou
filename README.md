# GO FANFOU

A Fanfou API Client SDK library for [Go](http://golang.org/)

[![Build Status](https://travis-ci.org/mogita/go-fanfou.svg?branch=master)](https://travis-ci.org/mogita/go-fanfou)
[![Coverage Status](https://coveralls.io/repos/github/mogita/go-fanfou/badge.svg?branch=master&service=github)](https://coveralls.io/github/mogita/go-fanfou?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mogita/go-fanfou)](https://goreportcard.com/report/github.com/mogita/go-fanfou)
[![GoDoc](https://godoc.org/github.com/mogita/go-fanfou?status.svg)](https://godoc.org/github.com/mogita/go-fanfou/fanfou)

**Please note that the library is at a very early stage of development.** Things could be changing at times. Breaking changes are expected, but I'll make it as less as possible. The stable versions will begin with the first [release](https://github.com/mogita/go-fanfou/releases) in the future.

## Usage

```
go get -u github.com/mogita/go-fanfou
```

```go
package main

import "github.com/mogita/go-fanfou/fanfou"

func main() {
  // ...
}
```

Please refer to the `examples` folder for the basic usages of this library.

Before running the examples, please fill the constants in `def.go` with your API keys and such. For obtaining a new API key please refer to https://fanfou.com/apps

You can run the examples to see how this library works:

```
go run examples/oauth/oauth.go
go run examples/xauth/xauth.go
go run examples/upload_photo/upload_photo.go
```

## Built With

* [oauth](https://github.com/mrjones/oauth) - OAuth 1.0 implementation in go (golang). http://www.mrjon.es

## Contributing

First of all, thank you very much for paying attention to this library. If you feel like helping improve it, please kindly make sure to follow the instructions:

Please link the pre-commit hook which runs tests and go-fmt before committing

```
ln -s $PWD/pre-commit.sh .git/hooks/pre-commit
```

Please always run tests before committing

```
go test ./...
```

## License

MIT © [mogita](https://github.com/mogita)
