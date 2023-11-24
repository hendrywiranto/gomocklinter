package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	reportMsg        = "calling Finish on gomock.Controller is no longer needed"
	pkgGolangMock    = "github.com"
	subPkgGolangMock = "golang"
	pkgUberMock      = "go.uber.org"
	mock             = "mock"
	controller       = "gomock.Controller"
	golangPkgLen     = 4
	uberPkgLen       = 3
	finish           = "Finish"
)

// New returns new gomocklinter analyzer.
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "gomocklinter",
		Doc:      "Checks the usage of go mocking libraries",
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
			// Get the token.File associated with the callExpr
			f := pass.Fset.File(callExpr.Pos())

			// Get the line number of the callExpr
			line := f.Line(callExpr.Pos())

			// Get the start position of the line
			lineStartPos := f.LineStart(line)

			// Get the start position of the next line
			nextLineStartPos := f.LineStart(line + 1)

			// Calculate the end position of the line
			lineEndPos := nextLineStartPos - 1

			// Create a fix to delete the line
			fix := analysis.SuggestedFix{
				Message: "Remove unnecessary call to Finish",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     lineStartPos,
						End:     lineEndPos,
						NewText: []byte(""),
					},
				},
			}

			// Add the fix to the diagnostic report
			diagnostic := analysis.Diagnostic{
				Pos:            callExpr.Pos(),
				Message:        reportMsg,
				SuggestedFixes: []analysis.SuggestedFix{fix},
			}
			pass.Report(diagnostic)
		}
	})

	return nil, nil
}

// isValidType checks whether t is a valid package source for gomock controller or not
// currently supports github.com/golang/mock/gomock.Controller and go.uber.org/mock/gomock.Controller.
func isValidType(t string) bool {
	if t[0] != '*' {
		return false
	}
	t = t[1:]

	strs := strings.Split(t, "/")
	return isTypeGomock(strs) || isTypeUberMock(strs)
}

// isTypeGomock checks if the given string slice represents gomock type from github.com/golang/mock/gomock.
// It returns true if the last four elements of the slice match the expected values for a gomock type.
// Otherwise, it returns false.
func isTypeGomock(strs []string) bool {
	if len(strs) < golangPkgLen {
		return false
	}
	strs = strs[len(strs)-golangPkgLen:]

	return strs[0] == pkgGolangMock && strs[1] == subPkgGolangMock && strs[2] == mock && strs[3] == controller
}

// isTypeUberMock checks if the given string slice represents gomock type from go.uber.org/mock/gomock.
// It returns true if the last three elements of the slice match the expected values for an Uber gomock type.
// Otherwise, it returns false.
func isTypeUberMock(strs []string) bool {
	if len(strs) < uberPkgLen {
		return false
	}
	strs = strs[len(strs)-uberPkgLen:]

	return strs[0] == pkgUberMock && strs[1] == mock && strs[2] == controller
}
