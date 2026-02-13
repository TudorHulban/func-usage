package funcusage

import (
	"fmt"
	"sort"
	"strings"
)

type AnalysisGroupedByPackage map[namePackage]LevelFunction

func (a AnalysisGroupedByPackage) PrintWith(printer *Printer) {
	pkgNames := make([]string, 0, len(a))

	for pkgName := range a {
		pkgNames = append(pkgNames, string(pkgName))
	}

	sort.Strings(pkgNames)

	for _, name := range pkgNames {
		fmt.Printf(
			"\n=== Package: %s (%d functions) ===\n",
			name,
			len(a[namePackage(name)]),
		)

		packageFunctions := a[namePackage(name)].OrderByNameAsc()

		packageFunctions.PrintWith(printer)
	}
}

func (a LevelFunction) GroupedByPackage() AnalysisGroupedByPackage {
	result := make(map[namePackage]LevelFunction)

	for _, fa := range a {
		pkg, errPackage := fa.getPackage()
		if errPackage != nil {
			fmt.Println(
				"GroupedByPackage:",
				errPackage,
			)

			continue
		}

		result[pkg] = append(result[pkg], fa)
	}

	return result
}

type AnalysisGroupedByObject map[nameObject]LevelFunction

func (a AnalysisGroupedByObject) PrintWith(printer *Printer) {
	objectNames := make([]string, 0, len(a))

	for objectName := range a {
		if len(objectName) == 0 {
			continue
		}

		objectNames = append(objectNames, string(objectName))
	}

	sort.Strings(objectNames)

	for _, name := range objectNames {
		fmt.Printf(
			"\n=== Object: %s (%d functions) ===\n",
			name,
			len(a[nameObject(name)]),
		)

		packageFunctions := a[nameObject(name)].OrderByNameAsc()

		packageFunctions.PrintWith(printer)
	}
}

func (a LevelFunction) GroupedByObject() AnalysisGroupedByObject {
	result := make(map[nameObject]LevelFunction)

	for _, fa := range a {
		result[fa.MethodOf] = append(result[fa.MethodOf], fa)
	}

	return result
}

type AnalysisGroupedByPackageAndObject map[namePackage]map[nameObject]LevelFunction

func (a AnalysisGroupedByPackageAndObject) PrintWith(printer *Printer) {
	packageNames := make([]string, 0, len(a))
	for pkg := range a {
		packageNames = append(packageNames, string(pkg))
	}

	sort.Strings(packageNames)

	for _, pkgName := range packageNames {
		fmt.Printf(
			"\n=== Package: %s ===\n",
			pkgName,
		)

		objMap := a[namePackage(pkgName)]
		objectNames := make([]string, 0, len(objMap))

		for obj := range objMap {
			objectNames = append(objectNames, string(obj))
		}

		sort.Strings(objectNames)

		for _, objName := range objectNames {
			fmt.Printf(
				"\n(object) %s (%d):\n",

				objName,
				len(objMap[nameObject(objName)]),
			)

			objMap[nameObject(objName)].PrintWith(printer)
		}
	}
}

func (a LevelFunction) GroupedByPackageAndObject() AnalysisGroupedByPackageAndObject {
	result := make(map[namePackage]map[nameObject]LevelFunction)

	for _, fa := range a {
		keyPackage, errPackage := fa.getPackage()
		if errPackage != nil {
			fmt.Println(
				"GroupedByPackageAndObject:",
				errPackage,
			)

			continue
		}

		keyObject := fa.MethodOf
		if len(keyObject) == 0 {
			keyObject = _LabelFunction
		}

		// Initialize package map if needed
		if result[keyPackage] == nil {
			result[keyPackage] = make(map[nameObject]LevelFunction)
		}

		// Append to object group
		result[keyPackage][keyObject] = append(
			result[keyPackage][keyObject],
			fa,
		)
	}

	return result
}

type AnalysisGroupedBySignature map[string]LevelFunction

func (a AnalysisGroupedBySignature) PrintWith(printer *Printer) {
	signatures := make([]string, 0, len(a))

	for signature := range a {
		if len(signature) == 0 {
			continue
		}

		signatures = append(signatures, signature)
	}

	sort.Strings(signatures)

	for _, signature := range signatures {
		fmt.Printf(
			"\n=== Signature: %s (%d functions) ===\n",
			signature,
			len(a[signature]),
		)

		group := a[signature].OrderByNameAsc()

		group.PrintWith(printer)
	}
}

func (a LevelFunction) GroupedByParamSignature() AnalysisGroupedBySignature {
	result := make(AnalysisGroupedBySignature, len(a))

	for _, usage := range a {
		key := strings.Join(usage.TypesParams, ",")

		group := result[key]
		group = append(group, usage)
		result[key] = group
	}

	return result
}

func (a LevelFunction) GroupedByResultSignature() AnalysisGroupedBySignature {
	result := make(AnalysisGroupedBySignature, len(a))

	for _, usage := range a {
		key := strings.Join(usage.TypesResults, ",")

		group := result[key]
		group = append(group, usage)
		result[key] = group
	}

	return result
}
