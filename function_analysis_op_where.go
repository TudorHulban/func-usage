package funcusage

import "fmt"

func (a Analysis) WhereNameIs(name string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.Name == name {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WherePackageIs(name string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		packageName, errPackage := usage.getPackage()
		if errPackage != nil {
			fmt.Println(
				"WherePackageIs:",
				errPackage,
			)

			continue
		}

		if packageName == namePackage(name) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereUnused() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.InternalCount == 0 && usage.ExternalCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereExported() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereUnexported() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if !isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereNotTested() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.InternalTestsCount == 0 && usage.ExternalTestsCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereTestedInternally() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.InternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WhereTestedExternally() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.ExternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}
