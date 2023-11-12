package examples

type testDummy string

func (td *testDummy) Finish() {}

func New() testDummy {
	return testDummy("this is a test")
}

func Finish() {
	td := New()
	td.Finish()
}
