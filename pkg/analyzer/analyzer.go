package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	finish     = "Finish"
	reportMsg  = "calling Finish on gomock.Controller is no longer needed"
	mock       = "mock"
	controller = "gomock.Controller"
)

var (
	pkgSourcesMap = map[string]bool{
		"golang":      true,
		"go.uber.org": true,
	}
)

// New returns new gomockcontrollerfinish analyzer.
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "gomockcontrollerfinish",
		Doc:      "Checks whether an unnecessary call to .Finish() on gomock.Controller exists",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		selIdent, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return
		}

		// check for unnecessary call to gomock.Controller.Finish()
		if isValidType(pass.TypesInfo.TypeOf(selIdent).String()) && selectorExpr.Sel.Name == finish {
			pass.Reportf(callExpr.Pos(), reportMsg)
		}
	})

	return nil, nil
}

// isValidType checks whether t is a valid package source for gomock controller or not
// currently supports golang/mock/gomock.Controller and go.uber.org/mock/gomock.Controller
//
// value of t can be *examples/vendor/go.uber.org/mock/gomock.Controller
// hence the checking is only the last 3 part
func isValidType(t string) bool {
	strs := strings.Split(t, "/")

	if len(strs) < 3 {
		return false
	}

	// get the last 3 elements
	strs = strs[len(strs)-3:]

	// first element has to be either golang or go.uber.org
	// second element has to be mock
	// third element has to be gomock.Controller
	return pkgSourcesMap[strs[0]] && strs[1] == mock && strs[2] == controller
}
