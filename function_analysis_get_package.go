package funcusage

import (
	"fmt"
	"strings"
)

func (fa *AnalysisFunction) getPackage() (namePackage, error) {
	// Case 1: filesystem-style path
	indexSlash := strings.LastIndex(fa.Key, "/")
	if indexSlash != -1 {
		rest := fa.Key[indexSlash+1:]

		pkg, _, foundSeparator := strings.Cut(rest, ".")
		if !foundSeparator {
			return "", fmt.Errorf(
				"invalid package path (missing . after /) in: %q",
				fa.Key,
			)
		}

		pkg = strings.ReplaceAll(pkg, "-", "")

		return namePackage(pkg), nil
	}

	// Case 2: package-qualified identifier
	indexDot := strings.Index(fa.Key, ".")
	if indexDot == -1 {
		return "", fmt.Errorf(
			"invalid package path (missing .) in: %q",
			fa.Key,
		)
	}

	pkg := strings.ReplaceAll(fa.Key[:indexDot], "-", "")

	return namePackage(pkg), nil
}
