package funcusage

import (
	"errors"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

func ensureNoPackageErrors(pkgs []*packages.Package) error {
	for _, packageFound := range pkgs {
		if len(packageFound.Errors) > 0 {
			for _, e := range packageFound.Errors {
				fmt.Printf(
					"Package %s error: %v\n",

					packageFound.PkgPath,
					e,
				)
			}

			return errors.New("packages have errors")
		}
	}

	return nil
}

func identOfCall(call *ast.CallExpr) *ast.Ident {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		return fun
	case *ast.SelectorExpr:
		return fun.Sel

	default:
		return nil
	}
}
