package examples_test

import (
	"testing"

	"examples/gomock" // gomock imitation
)

// This file tests that the linter won't be tricked by another package source that has the same Controller object as the original package.

func TestSimilarFinishCall(t *testing.T) {
	mock := gomock.NewController(t)
	mock.Finish()
}

func TestSimilarFinishCallDefer(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
}

func TestSimilarFinishCallWithoutT(t *testing.T) {
	mock := gomock.NewController(nil)
	mock.Finish()
}

func TestSimilarFinsihCallInAnotherFunction(t *testing.T) {
	mock := gomock.NewController(t)
	callSimilarFinish(mock)
}

func callSimilarFinish(mock *gomock.Controller) {
	mock.Finish()
}

func TestSimilarNoFinishCall(t *testing.T) {
	gomock.NewController(t)
}
