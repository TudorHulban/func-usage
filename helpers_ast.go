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

func extractMethodOf(fn *types.Func) NameObject {
	signature := fn.Type().(*types.Signature)

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
		return NameObject(
			named.Obj().Name(),
		)
	}

	return ""
}
