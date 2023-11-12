package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	gomockControllerType = "mock/gomock.Controller"
	finish               = "Finish"
	reportMsg            = "since go1.14, if you are passing a testing.T to NewController then calling Finish on gomock.Controller is no longer needed"
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

		// check if it's a test file
		if !isTestFile(pass.Fset.Position(callExpr.Pos()).Filename) {
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
		if strings.HasSuffix(pass.TypesInfo.TypeOf(selIdent).String(), gomockControllerType) && selectorExpr.Sel.Name == finish {
			pass.Reportf(selectorExpr.Sel.Pos(), reportMsg)
		}
	})

	return nil, nil
}

// isTestFile checks if the file is a test file based on its name.
func isTestFile(filename string) bool {
	return strings.HasSuffix(filename, "_test.go")
}
