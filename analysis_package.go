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

// StatisticsForPackage reports metrics only for functions defined in this package.
//
// Notes:
//   - PackageFunctions contains only functions declared in the real package
//     (no *_test.go, no external test packages, no synthetic test packages).
//   - Global analysis (LevelFunction) also includes:
//   - functions from *_test.go
//   - functions from external test packages (package xxx_test)
//   - synthetic test/binary functions (init, main, helpers)
//   - Therefore, counts in package statistics (functions, methods, untested,
//     unused, etc.) can legitimately differ from the global analysis numbers.
//     Package-level stats are the "production/package view", while LevelFunction
//     is the "full view including tests".
func (level LevelPackage) StatisticsForPackage(forModulePath, pkgName string) string {
	var b strings.Builder

	for _, pkg := range level {
		if pkg.Name != pkgName {
			continue
		}

		methodCount := 0
		methodsPerObject := make(map[string]int)

		untested := 0
		unused := 0

		internalCalls := 0
		externalCalls := 0

		for _, fa := range pkg.PackageFunctions {
			if len(fa.MethodOf) != 0 {
				methodCount++
				methodsPerObject[string(fa.MethodOf)]++
			}

			// untested
			if fa.InternalTestsCount+fa.ExternalTestsCount == 0 {
				untested++
			}

			// unused
			totalCalls := fa.InternalCount +
				fa.ExternalCount +
				fa.InternalTestsCount +
				fa.ExternalTestsCount

			if totalCalls == 0 {
				unused++
			}

			// cohesion
			internalCalls = internalCalls + fa.InternalCount
			externalCalls = externalCalls + fa.ExternalCount
		}

		// avg methods per object
		var avgMethods float64
		if len(methodsPerObject) > 0 {
			avgMethods = float64(methodCount) / float64(len(methodsPerObject))
		}

		var (
			maxMethods int
			maxObject  string
		)

		for obj, count := range methodsPerObject {
			if count > maxMethods {
				maxMethods = count
				maxObject = obj
			}
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

		fmt.Fprintf(&b, "Functions and Methods: %d\n", len(pkg.PackageFunctions))
		fmt.Fprintf(&b, "Methods: %d\n", methodCount)

		fmt.Fprintf(&b, "Types: %d\n", len(pkg.Types))
		fmt.Fprintf(&b, "Avg methods per object: %.2f\n", avgMethods)
		fmt.Fprintf(&b, "Max methods per object: %d (%s)\n", maxMethods, maxObject)

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

// Package statistics intentionally exclude test-only functions.
// PackageFunctions contains only functions declared in the real package
// (no *_test.go, no external test packages, no synthetic test packages).
// Global analysis (LevelFunction) includes all of these.
// Therefore, comparing package statistics with the global function list
// will show differences. This is by design, not a bug.
func (level LevelPackage) Statistics(forModulePath string) string {
	var b strings.Builder

	for _, pkg := range level {
		b.WriteString(
			level.StatisticsForPackage(forModulePath, pkg.Name),
		)
	}

	return b.String()
}
