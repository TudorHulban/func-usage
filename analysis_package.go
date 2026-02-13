package funcusage

import (
	"fmt"
	"strings"
)

type inUse bool

type AnalysisPackage struct {
	Name             string
	Types            map[string]inUse
	PackageFunctions LevelFunction
}

type LevelPackage []AnalysisPackage

func (level LevelPackage) String() string {
	lines := make([]string, 0, len(level)+1)

	lines = append(lines, "Package,Functions,Methods,Types,TypesInUse,TypesUnused")

	for _, pa := range level {
		var funcCount, methodCount int

		for _, fn := range pa.PackageFunctions {
			if fn.MethodOf == "" {
				funcCount++
			} else {
				methodCount++
			}
		}

		totalTypes := len(pa.Types)

		var inUseCount int

		for _, used := range pa.Types {
			if used {
				inUseCount++
			}
		}

		unusedCount := totalTypes - inUseCount

		lines = append(
			lines,
			fmt.Sprintf(
				"%s,%d,%d,%d,%d,%d",

				pa.Name,
				funcCount,
				methodCount,
				totalTypes,
				inUseCount,
				unusedCount,
			),
		)
	}

	return strings.Join(lines, "\n")
}
