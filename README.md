# Go AWS Lambda Utils

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

## Overview

`awslambda` is helper package for AWS Lambda on Go.

See the docs on [pkg.go.dev](https://pkg.go.dev/github.com/grokify/go-awslambda).

## Installation

```bash
$ go get github.com/grokify/go-awslambda/...
```

## Usage

* `NewHTTPRequest()` can be used to create an `*http.Request` given an `events.APIGatewayProxyRequest`.
* `NewReaderMultipart()` can be used to create an `*multipart.Reader` given an `events.APIGatewayProxyRequest`.

See the multipart example here:

[`examples/multipart/main.go`](examples/multipart/main.go)

A description is available on Stack Overflow here:

https://stackoverflow.com/a/68496889/1908967

## More Info

1. [amazon.com: Set up Lambda proxy integrations in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-lambda-proxy-integrations.html#api-gateway-simple-proxy-for-lambda-input-format)

## Contributing

Features, Issues, and Pull Requests are always welcome.

To contribute:

1. Fork it ( http://github.com/grokify/go-awslambda/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

Please report issues and feature requests on [Github](https://github.com/grokify/go-awslambda).

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/go-awslambda/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/go-awslambda?badge
 [build-status-svg]: https://github.com/grokify/gp-awslambda/actions/workflows/ci.yaml/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/go-awslambda/actions/workflows/ci.yaml
 [build-status-svg]: https://github.com/grokify/gp-awslambda/actions/workflows/lint.yaml/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/go-awslambda/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-awslambda
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-awslambda
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-awslambda
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-awslambda
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-awslambda/blob/master/LICENSE
