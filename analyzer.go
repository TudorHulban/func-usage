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
	modulePath string
}

func NewAnalyzer(atRoot string) (*Analyzer, error) {
	modulePath, errGetModule := getModulePath(atRoot)
	if errGetModule != nil {
		return nil,
			errGetModule
	}

	return &Analyzer{
			root:       atRoot,
			modulePath: modulePath,
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

func (a Analyzer) Analyze(inMode AnalyzeMode, includeExternal bool) (Usage, error) { //nolint:revive,cognitive-complexity
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

	usages := make(map[string]*FunctionUsage)

	for _, packageFound := range packagesLoaded {
		callerPkg := packageFound.ID

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

						if fnName == "main" || fnName == "init" || strings.HasPrefix(fnName, "Test") {
							return true
						}

						definingPkg := fn.Pkg().Path()

						isOutsideModule := !strings.HasPrefix(definingPkg, a.modulePath)
						if isOutsideModule && !includeExternal {
							return true
						}

						key := buildFuncKey(fn)

						usage := usages[key]
						if usage == nil {
							usage = &FunctionUsage{
								Key:      key,
								Name:     fnName,
								Position: packageFound.Fset.Position(fn.Pos()),
							}

							usages[key] = usage
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

					calledPkg := fn.Pkg().Path()

					isOutsideModule := !strings.HasPrefix(calledPkg, a.modulePath)
					if isOutsideModule && !includeExternal {
						return true
					}

					key := buildFuncKey(fn)

					usage := usages[key]
					if usage == nil {
						usage = &FunctionUsage{
							Key:      key,
							Name:     fn.Name(),
							Position: packageFound.Fset.Position(fn.Pos()),
						}

						usages[key] = usage
					}

					usage.updateOccurences(callerPkg, calledPkg, isTest)

					return true
				},
			)
		}
	}

	result := make([]FunctionUsage, 0, len(usages))

	for _, usage := range usages {
		result = append(result, *usage)
	}

	return result,
		nil
}
