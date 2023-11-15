# gomockcontrollerfinish

![CI](https://github.com/hendrywiranto/gomockcontrollerfinish/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/hendrywiranto/gomockcontrollerfinish)](https://goreportcard.com/report/github.com/hendrywiranto/gomockcontrollerfinish)
[![Coverage Status](https://coveralls.io/repos/github/hendrywiranto/gomockcontrollerfinish/badge.svg?branch=main)](https://coveralls.io/github/hendrywiranto/gomockcontrollerfinish?branch=main)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A linter that checks whether an unnecessary call to .Finish() on gomock.Controller exists

Note: The original [golang/mock](https://github.com/golang/mock) package is archived and the maintained fork is [go.uber.org/mock](https://github.com/uber/mock). This linter supports both.

## Installation & usage

```
$ go install github.com/hendrywiranto/gomockcontrollerfinish@latest
$ gomockcontrollerfinish ./...
```

or build the binary and use `go vet`
```
$ go build -o gomockcontrollerfinish main.go
$ go vet -vettool=./gomockcontrollerfinish ./...
```

## Motivation
As highlighted in https://pkg.go.dev/github.com/golang/mock/gomock#NewController

>New in go1.14+, if you are passing a *testing.T into this function you no longer need to call ctrl.Finish() in your test methods.  

## Examples
```
// Bad
func TestFoo(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // no need to call this since go1.14
}

// Good
func TestFoo(t *testing.T) {
	mock := gomock.NewController(t)	
}
```
