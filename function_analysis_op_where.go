package funcusage

import (
	"fmt"
	"slices"
)

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

func (a Analysis) WithNoParams() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if len(usage.TypesParams) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WithNoResults() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if len(usage.TypesResults) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) WithErrorReturn() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if slices.Contains(usage.TypesResults, "error") {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) AcceptingOnly(typeNames ...string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if unorderedButSameItems(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) ReturningOnly(typeNames ...string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if unorderedButSameItems(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) Accepting(typeNames ...string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if unorderedButContainsAll(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) Returning(typeNames ...string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if unorderedButContainsAll(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}
