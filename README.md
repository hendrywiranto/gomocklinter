# gomocklinter

![CI](https://github.com/hendrywiranto/gomocklinter/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/hendrywiranto/gomocklinter)](https://goreportcard.com/report/github.com/hendrywiranto/gomocklinter)
[![Coverage Status](https://coveralls.io/repos/github/hendrywiranto/gomocklinter/badge.svg?branch=main)](https://coveralls.io/github/hendrywiranto/gomocklinter?branch=main)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

A linter that checks the usage of go mocking libraries

Note: The original [golang/mock](https://github.com/golang/mock) package is archived and the maintained fork is [go.uber.org/mock](https://github.com/uber/mock).  
This linter supports both.

## Installation & usage

```
$ go install github.com/hendrywiranto/gomocklinter@latest
$ gomocklinter ./...
```

or build the binary and use `go vet`
```
$ go build -o gomocklinter main.go
$ go vet -vettool=./gomocklinter ./...
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
