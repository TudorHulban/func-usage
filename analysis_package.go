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

func (level LevelPackage) Statistics(forModulePath string) string {
	var b strings.Builder

	for _, pkg := range level {
		funcCount := len(pkg.PackageFunctions)

		methodCount := 0
		methodsPerReceiver := make(map[string]int)

		untested := 0
		unused := 0

		internalCalls := 0
		externalCalls := 0

		for _, fn := range pkg.PackageFunctions {
			// methods
			if fn.MethodOf != "" {
				methodCount++
				methodsPerReceiver[string(fn.MethodOf)]++
			}

			// untested
			if fn.InternalTestsCount+fn.ExternalTestsCount == 0 {
				untested++
			}

			// unused
			totalCalls := fn.InternalCount +
				fn.ExternalCount +
				fn.InternalTestsCount +
				fn.ExternalTestsCount

			if totalCalls == 0 {
				unused++
			}

			// cohesion
			internalCalls += fn.InternalCount
			externalCalls += fn.ExternalCount
		}

		// avg methods per object
		avgMethods := 0.0
		if len(methodsPerReceiver) > 0 {
			avgMethods = float64(methodCount) / float64(len(methodsPerReceiver))
		}

		// cohesion score
		cohesion := 1.0
		if internalCalls+externalCalls > 0 {
			cohesion = float64(internalCalls) / float64(internalCalls+externalCalls)
		}

		// unused types
		unusedTypes := make([]string, 0)

		for t, used := range pkg.Types {
			base := t
			for strings.HasPrefix(base, "*") {
				base = base[1:]
			}

			if !bool(used) && isRelevantType(t, forModulePath) {
				unusedTypes = append(unusedTypes, base)
			}
		}

		// print
		fmt.Fprintf(&b, "Package: %s\n", pkg.Name)
		fmt.Fprintf(&b, "Functions: %d\n", funcCount)
		fmt.Fprintf(&b, "Methods: %d\n", methodCount)
		fmt.Fprintf(&b, "Avg methods per object: %.2f\n", avgMethods)
		fmt.Fprintf(&b, "Types: %d\n", len(pkg.Types))

		fmt.Fprintf(&b, "Untested functions: %d\n", untested)
		fmt.Fprintf(&b, "Unused functions: %d\n", unused)
		fmt.Fprintf(&b, "Cohesion: %.2f\n", cohesion)

		fmt.Fprintf(&b, "Unused types: %d\n", len(unusedTypes))

		if len(unusedTypes) > 0 {
			for i, t := range unusedTypes {
				fmt.Fprintf(&b, "%d. %s\n", i+1, t)
			}
		}
	}

	return b.String()
}
