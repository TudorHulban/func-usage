package funcusage

import (
	"fmt"
	"sort"
)

type AnalysisGroupedByPackage map[NamePackage]Analysis

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
			len(a[NamePackage(name)]),
		)

		packageFunctions := a[NamePackage(name)].OrderByNameAsc()

		packageFunctions.PrintWith(printer)
	}
}

func (a Analysis) GroupedByPackage() AnalysisGroupedByPackage {
	result := make(map[NamePackage]Analysis)

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

type AnalysisGroupedByObject map[NameObject]Analysis

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
			len(a[NameObject(name)]),
		)

		packageFunctions := a[NameObject(name)].OrderByNameAsc()

		packageFunctions.PrintWith(printer)
	}
}

func (a Analysis) GroupedByObject() AnalysisGroupedByObject {
	result := make(map[NameObject]Analysis)

	for _, fa := range a {
		result[fa.MethodOf] = append(result[fa.MethodOf], fa)
	}

	return result
}

type AnalysisGroupedByPackageAndObject map[NamePackage]map[NameObject]Analysis

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

		objMap := a[NamePackage(pkgName)]
		objectNames := make([]string, 0, len(objMap))

		for obj := range objMap {
			objectNames = append(objectNames, string(obj))
		}
		sort.Strings(objectNames)

		for _, objName := range objectNames {
			fmt.Printf(
				"\n(object) %s (%d):\n",

				objName,
				len(objMap[NameObject(objName)]),
			)

			objMap[NameObject(objName)].PrintWith(printer)
		}
	}
}

func (a Analysis) GroupedByPackageAndObject() AnalysisGroupedByPackageAndObject {
	result := make(map[NamePackage]map[NameObject]Analysis)

	for _, fa := range a {
		pkg, errPackage := fa.getPackage()
		if errPackage != nil {
			fmt.Println(
				"GroupedByPackageAndObject:",
				errPackage,
			)

			continue
		}

		keyPackage := NamePackage(pkg)

		keyObject := fa.MethodOf
		if len(keyObject) == 0 {
			keyObject = _LabelFunction
		}

		// Initialize package map if needed
		if result[keyPackage] == nil {
			result[keyPackage] = make(map[NameObject]Analysis)
		}

		// Append to object group
		result[keyPackage][keyObject] = append(
			result[keyPackage][keyObject],
			fa,
		)
	}

	return result
}
