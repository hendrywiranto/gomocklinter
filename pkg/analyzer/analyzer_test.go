package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gostaticanalysis/testutil"
	"github.com/hendrywiranto/gomocklinter/pkg/analyzer"
)

func TestGoMockLinter(t *testing.T) {
	pkgs := []string{
		"examples",
	}
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, analyzer.New(), pkgs...)
}
