package examples

import (
	"testing"

	"github.com/golang/mock/gomock"
	gomick "github.com/golang/mock/gomock"
)

func TestRenamedFinishCall(t *testing.T) {
	mock := gomick.NewController(t)
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedFinishCallDefer(t *testing.T) {
	mock := gomick.NewController(t)
	defer mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedNoFinishCall(t *testing.T) {
	gomick.NewController(t)
}

func TestRenamedFinishCallWithoutT(t *testing.T) {
	mock := gomick.NewController(nil)
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}

func TestRenamedFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomick.NewController(t)
	renamedCallFinish(mock)
}

func renamedCallFinish(mock *gomock.Controller) {
	mock.Finish() // want "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
}
