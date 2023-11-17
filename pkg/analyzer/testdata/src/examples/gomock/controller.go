package gomock

// TestReporter is something that can be used to report test failures.
// It imitates TestReporter from original gomock package.
type TestReporter interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// Controller imitates the original Controller from gomock package
type Controller struct{}

// Controller imitates the original Finish() on Controller from gomock package
func (td *Controller) Finish() {}

// NewController imitates the original NewController from gomock package.
// It returns Controller imitation.
func NewController(t TestReporter) *Controller {
	return &Controller{}
}
