package examples

type testDummy struct{}

func (td *testDummy) Finish() {}

func New() *testDummy {
	return &testDummy{}
}

// finish call from the package that defines it
func Finish() {
	td := New()
	td.Finish()
}
