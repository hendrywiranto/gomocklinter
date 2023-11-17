package examples_test

import (
	"testing"

	gomick "go.uber.org/mock/gomock"
)

// This file tests that the linter will still work even if the gomock that is from the go.uber.org fork is renamed to something else

func TestUberRenamedFinishCall(t *testing.T) {
	mock := gomick.NewController(t)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberRenamedFinishCallDefer(t *testing.T) {
	mock := gomick.NewController(t)
	defer mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberRenamedFinishCallWithoutT(t *testing.T) {
	mock := gomick.NewController(nil)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberRenamedFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomick.NewController(t)
	uberRenamedCallFinish(mock)
}

func uberRenamedCallFinish(mock *gomick.Controller) {
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberRenamedNoFinishCall(t *testing.T) {
	gomick.NewController(t)
}
