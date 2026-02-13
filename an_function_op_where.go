package funcusage

import (
	"fmt"
)

func (a LevelFunction) WhereNameIs(name string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.Name == name {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WherePackageIs(name string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

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

func (a LevelFunction) WhereUnused() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.InternalCount == 0 && usage.ExternalCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WhereExported() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WhereUnexported() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if !isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WhereNotTested() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.InternalTestsCount == 0 && usage.ExternalTestsCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WhereTestedInternally() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.InternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WhereTestedExternally() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.ExternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}
