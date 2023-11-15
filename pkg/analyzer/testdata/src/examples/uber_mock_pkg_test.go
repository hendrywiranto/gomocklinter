package examples_test

import (
	"examples"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUberFinishCall(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberFinishCallDefer(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberFinishCallWithoutT(t *testing.T) {
	mock := gomock.NewController(nil)
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomock.NewController(t)
	callUberFinish(mock)
}

func callUberFinish(mock *gomock.Controller) {
	mock.Finish() // want "calling Finish on gomock.Controller is no longer needed"
}

func TestUberNoFinishCall(t *testing.T) {
	gomock.NewController(t)
}

func TestUberFinishCallOther(t *testing.T) {
	mock := examples.New()
	mock.Finish()
}
