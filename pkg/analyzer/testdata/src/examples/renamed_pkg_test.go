package examples_test

import (
	"testing"

	gomick "github.com/golang/mock/gomock"
)

// This file tests that the linter will still work even if the gomock from the original golang package is renamed

func TestRenamedFinishCall(t *testing.T) {
	mock := gomick.NewController(t)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedFinishCallDefer(t *testing.T) {
	mock := gomick.NewController(t)
	defer mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedFinishCallWithoutT(t *testing.T) {
	mock := gomick.NewController(nil)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomick.NewController(t)
	renamedCallFinish(mock)
}

func renamedCallFinish(mock *gomick.Controller) {
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedNoFinishCall(t *testing.T) {
	gomick.NewController(t)
}
