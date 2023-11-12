package examples

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestFinishCall(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestFinishCallDefer(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestFinishCallWithoutT(t *testing.T) {
	mock := gomock.NewController(nil)
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomock.NewController(t)
	callFinish(mock)
}

func callFinish(mock *gomock.Controller) {
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestNoFinishCall(t *testing.T) {
	gomock.NewController(t)
}

func TestFinishCallOther(t *testing.T) {
	mock := New()
	mock.Finish()
}
