package funcusage

import (
	"strings"
)

// Empty string ("") → internal (same package)
// Stdlib (fmt, strings) → external (doesn't start with your module path)
// Your module (github.com/tudor/project/*) → internal (prefix match)
// External deps (github.com/other/module) → external (no prefix match)
// Local module (simple, simple/foo) → internal (prefix match)

func isExternalPackage(pkgPath, modulePath string) bool {
	// Empty path means same package (unexported functions)
	if len(pkgPath) == 0 {
		return false
	}

	// Internal: starts with module path
	if strings.HasPrefix(pkgPath, modulePath) {
		return false
	}

	// Everything else (stdlib + third-party) is external
	return true
}

// isExportedName reports whether a Go identifier is exported.
// A name is exported if it begins with an upper-case letter.
func isExportedName(name string) bool {
	if len(name) == 0 {
		return false
	}

	r := name[0]

	return r >= 'A' && r <= 'Z'
}

// isTestFile reports whether a filename belongs to a Go test file.
// Centralizing this avoids scattering string suffix checks.
func isTestFile(filename string) bool {
	return strings.HasSuffix(filename, "_test.go")
}

func isRelevantPackage(id string) bool {
	if strings.HasSuffix(id, ".test") {
		return true
	}

	if strings.Contains(id, " [") { // e.g. "pkg [pkg.test]"
		return true
	}

	if strings.HasSuffix(id, "_test") {
		return true
	}

	return false
}

func isRelevantType(t string, modulePath string) bool {
	// strip pointer
	for strings.HasPrefix(t, "*") {
		t = t[1:]
	}

	// strip slice
	for strings.HasPrefix(t, "[]") {
		t = t[2:]
	}

	// strip map key/value (very naive but good enough)
	if strings.HasPrefix(t, "map[") {
		return false
	}

	// keep only types from our module
	return strings.HasPrefix(t, modulePath+".")
}
