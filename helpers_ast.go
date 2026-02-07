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

func extractMethodOf(fnDecl *ast.FuncDecl, typeInfo *types.Info) NameObject {
	if fnDecl.Recv == nil || len(fnDecl.Recv.List) == 0 {
		return "" // Not a method
	}

	receiver := fnDecl.Recv.List[0]
	receiverType := typeInfo.TypeOf(receiver.Type)
	if receiverType == nil {
		return ""
	}

	typeStr := receiverType.String()

	// Remove package prefix: "*github.com/me/project/pkg.User" -> "*User"
	if lastDot := strings.LastIndex(typeStr, "."); lastDot != -1 {
		return NameObject(typeStr[lastDot+1:])
	}

	return NameObject(typeStr)
}
