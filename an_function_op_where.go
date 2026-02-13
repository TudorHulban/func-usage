package funcusage

import (
	"fmt"
)

func (level LevelFunction) WhereNameIs(name string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.Name == name {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WherePackageIs(name string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
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

func (level LevelFunction) WhereUnused() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.InternalCount == 0 && usage.ExternalCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WhereExported() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WhereUnexported() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if !isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WhereNotTested() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.InternalTestsCount == 0 && usage.ExternalTestsCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WhereTestedInternally() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.InternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WhereTestedExternally() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.ExternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}
