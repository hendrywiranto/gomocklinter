package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	gomockController = "github.com/golang/mock/gomock.Controller"
	gomock           = "gomock"
	finish           = "Finish"
	newController    = "NewController"
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

	// track if gomock.NewController(s.T()) is called in the current test file
	var newControllerCalled bool

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		// Check if it's a test file
		if !isTestFile(pass.Fset.Position(callExpr.Pos()).Filename) {
			return
		}

		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}

		if selIdent, ok := selectorExpr.X.(*ast.Ident); ok && selIdent.Name == gomock && selectorExpr.Sel.Name == newController {
			if len(callExpr.Args) == 1 {
				argExpr, ok := callExpr.Args[0].(*ast.SelectorExpr)
				if ok {
					// check if it's gomock.NewController(s.T())
					if tIdent, ok := argExpr.X.(*ast.Ident); ok && tIdent.Name == "s" && argExpr.Sel.Name == "T" {
						newControllerCalled = true
					}
				}
			}
		}

		// check for unnecessary call to gomock.Controller.Finish()
		if newControllerCalled && pass.TypesInfo.TypeOf(callExpr).String() == gomockController && selectorExpr.Sel.Name == finish {
			pass.Reportf(callExpr.Pos(), "since go1.14+, if you are passing a *testing.T to NewController() then calling Finish() on gomock.Controller is no longer needed")
		}
	})

	return nil, nil
}

// isTestFile checks if the file is a test file based on its name.
func isTestFile(filename string) bool {
	return strings.HasSuffix(filename, "_test.go")
}
