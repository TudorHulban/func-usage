package funcusage

import (
	"go/ast"
	"go/types"
	"strings"
)

// buildFuncKey constructs a canonical identity string for a function or method.
// For functions: "<pkgPath>.<Name>"
// For methods:   "<pkgPath>.<RecvType>.<Name>"
//
// The key is stable and deterministic, making it suitable for maps and sorting.
func buildFuncKey(fn *types.Func) string {
	if fn == nil || fn.Pkg() == nil {
		return ""
	}

	pkgPath := fn.Pkg().Path()

	sig, couldCast := fn.Type().(*types.Signature)
	if !couldCast {
		return ""
	}

	recv := sig.Recv()
	if recv == nil {
		return pkgPath + "." + fn.Name() // Plain function.
	}

	// Method: normalize receiver type to avoid spaces.
	recvType := strings.ReplaceAll(recv.Type().String(), " ", "")

	return pkgPath + "." + recvType + "." + fn.Name()
}

// extractCallExpr returns the call expression if n is a call, otherwise nil.
func extractCallExpr(n ast.Node) *ast.CallExpr {
	call, _ := n.(*ast.CallExpr) //nolint:revive,unchecked-type-assertion

	return call
}

func extractMethodOf(fn *types.Func) nameObject {
	signature, couldCast := fn.Type().(*types.Signature)
	if !couldCast {
		return ""
	}

	receiver := signature.Recv()
	if receiver == nil {
		return "" // Not a method
	}

	receiverType := receiver.Type()

	// Strip pointer if present
	if ptr, couldCaast := receiverType.(*types.Pointer); couldCaast {
		receiverType = ptr.Elem()
	}

	// Get the named type
	if named, couldCast := receiverType.(*types.Named); couldCast {
		return nameObject(
			named.Obj().Name(),
		)
	}

	return ""
}

func extractSignatureTypes(fn *types.Func) ([]string, []string) {
	signature, couldCast := fn.Type().(*types.Signature)
	if !couldCast {
		return nil, nil
	}

	qualifier := func(pkg *types.Package) string {
		if pkg == nil {
			return ""
		}

		return pkg.Path()
	}

	params := make([]string, 0, signature.Params().Len())

	for param := range signature.Params().Variables() {
		params = append(
			params,
			types.TypeString(param.Type(), qualifier),
		)
	}

	results := make([]string, 0, signature.Results().Len())

	for result := range signature.Results().Variables() {
		results = append(
			results,
			types.TypeString(result.Type(), qualifier),
		)
	}

	return params, results
}
