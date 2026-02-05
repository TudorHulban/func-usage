package funcusage

import (
	"errors"
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"path/filepath"
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

func getModulePath(root string) (string, error) {
	modPath := filepath.Join(root, "go.mod")

	data, err := os.ReadFile(modPath) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to read go.mod: %w", err)
	}

	lines := strings.Split(string(data), "\n") //nolint:modernize

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	return "",
		errors.New("module path not found in go.mod")
}
