package funcusage

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Analyzer struct {
	root       string
	ModulePath string

	DefinedTypes  []string
	ExportedTypes []string
}

func NewAnalyzer(atRoot string) (*Analyzer, error) {
	modulePath, errGetModule := getModulePath(atRoot)
	if errGetModule != nil {
		return nil,
			errGetModule
	}

	return &Analyzer{
			root:       atRoot,
			ModulePath: modulePath,
		},
		nil
}

func (a *Analyzer) loadPackages() ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo |
			packages.NeedFiles |
			packages.NeedForTest |
			packages.NeedModule,
		Dir:   a.root,
		Tests: true,
	}

	return packages.Load(cfg, "./...")
}

func (a Analyzer) Analyze(inMode AnalyzeMode, includeExternal bool) (*Analysis, error) { //nolint:gocyclo,revive,cognitive-complexity
	packagesLoaded, errLoad := a.loadPackages()
	if errLoad != nil {
		return nil,
			fmt.Errorf(
				"failed to load packages: %w",
				errLoad,
			)
	}

	if errPackageLoad := ensureNoPackageErrors(packagesLoaded); errPackageLoad != nil {
		return nil, errPackageLoad
	}

	usages := make(map[string]*AnalysisFunction)
	packagesMap := make(map[string]*AnalysisPackage)

	for _, packageFound := range packagesLoaded {
		packageFoundID := packageFound.ID

		var pkgEntry *AnalysisPackage

		if !isSyntheticTestPackage(packageFoundID) {
			pkgEntry = packagesMap[packageFoundID]
			if pkgEntry == nil {
				pkgEntry = &AnalysisPackage{
					Name:             packageFoundID,
					Types:            make(map[string]inUse),
					PackageFunctions: make(LevelFunction, 0, 32),
				}

				packagesMap[packageFoundID] = pkgEntry
			}
		}

		for ix, file := range packageFound.Syntax {
			if ix >= len(packageFound.GoFiles) { // checks the assumption that GoFiles match Syntax 1:1.
				continue
			}

			filename := packageFound.GoFiles[ix]
			isTest := isTestFile(filename)

			if skipFileByMode(isTest, inMode) {
				continue
			}

			ast.Inspect(
				file,
				func(n ast.Node) bool {
					if fnDeclaration, couldCast := n.(*ast.FuncDecl); couldCast {
						obj := packageFound.TypesInfo.Defs[fnDeclaration.Name]
						if obj == nil {
							return true // should not happen after type-check
						}

						fn, asFunction := obj.(*types.Func)
						if !asFunction {
							return true // e.g. blank func name _
						}

						fnName := fn.Name()

						if fnName == "main" ||
							fnName == "init" ||
							strings.HasPrefix(fnName, "Test") ||
							strings.HasPrefix(fnName, "Benchmark") ||
							strings.HasPrefix(fnName, "Fuzz") {
							return true
						}

						definingPkg := fn.Pkg().Path()

						isOutsideModule := !strings.HasPrefix(definingPkg, a.ModulePath)
						if isOutsideModule && !includeExternal {
							return true
						}

						key := buildFuncKey(fn)

						usage := usages[key]
						if usage == nil {
							usage = &AnalysisFunction{
								Key:      key,
								Name:     fnName,
								Position: packageFound.Fset.Position(fn.Pos()),
							}

							usage.MethodOf = extractMethodOf(fn)
							usage.TypesParams, usage.TypesResults = extractSignatureTypes(fn)

							usages[key] = usage

							// attach function to package
							if pkgEntry != nil {
								pkgEntry.PackageFunctions = append(pkgEntry.PackageFunctions, *usage)

								// register types in package
								for _, t := range usage.TypesParams {
									pkgEntry.Types[t] = false
								}

								for _, t := range usage.TypesResults {
									pkgEntry.Types[t] = false
								}
							}
						}

						return true
					}

					call := extractCallExpr(n)
					if call == nil {
						return true
					}

					obj := packageFound.TypesInfo.Uses[identOfCall(call)]

					fn, couldCast := obj.(*types.Func)
					if !couldCast {
						return true
					}

					pkgCandidate := fn.Pkg()
					if pkgCandidate == nil {
						return true
					}

					pkgCalled := pkgCandidate.Path()

					isOutsideModule := !strings.HasPrefix(pkgCalled, a.ModulePath)
					if isOutsideModule && !includeExternal {
						return true
					}

					key := buildFuncKey(fn)

					usage := usages[key]
					if usage == nil {
						usage = &AnalysisFunction{
							Key:      key,
							Name:     fn.Name(),
							Position: packageFound.Fset.Position(fn.Pos()),
						}

						usage.MethodOf = extractMethodOf(fn)
						usage.TypesParams, usage.TypesResults = extractSignatureTypes(fn)

						usages[key] = usage
					}

					usage.updateOccurences(packageFoundID, pkgCalled, isTest)

					return true
				},
			)
		}
	}

	resultAnalysisFunction := make([]AnalysisFunction, 0, len(usages))

	for _, usage := range usages {
		resultAnalysisFunction = append(resultAnalysisFunction, *usage)
	}

	resultPackages := make(LevelPackage, 0, len(packagesMap))

	for _, pkg := range packagesMap {
		resultPackages = append(resultPackages, *pkg)
	}

	return &Analysis{
			LevelFunction: resultAnalysisFunction,
			LevelPackage:  resultPackages,
		},
		nil
}
