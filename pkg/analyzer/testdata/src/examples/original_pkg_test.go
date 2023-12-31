package examples_test

import (
	"testing"

	"examples"

	"github.com/golang/mock/gomock"
)

// This file tests that the linter will still work if the gomock is from the original golang package

func TestFinishCall(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestFinishCallDefer(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestFinishCallWithoutT(t *testing.T) {
	mock := gomock.NewController(nil)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomock.NewController(t)
	callFinish(mock)
}

func callFinish(mock *gomock.Controller) {
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestNoFinishCall(t *testing.T) {
	gomock.NewController(t)
}

func TestFinishCallOther(t *testing.T) {
	mock := examples.New()
	mock.Finish()
}
