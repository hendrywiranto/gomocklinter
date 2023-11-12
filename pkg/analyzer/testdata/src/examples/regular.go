package examples

type testDummy struct{}

func (td *testDummy) Finish() {}

func New() *testDummy {
	return &testDummy{}
}

func Finish() {
	td := New()
	td.Finish()
}
