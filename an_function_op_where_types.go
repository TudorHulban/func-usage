package funcusage

import (
	"slices"
	"strings"
)

func (level LevelFunction) WithNoParams() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if len(fa.TypesParams) == 0 {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) WithNoResults() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if len(fa.TypesResults) == 0 {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) WithErrorReturn() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if slices.Contains(fa.TypesResults, "error") {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) AcceptingOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if unorderedButSameItems(fa.TypesParams, typeNames) {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) ReturningOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if unorderedButSameItems(fa.TypesResults, typeNames) {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) Accepting(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if unorderedButContainsAll(fa.TypesParams, typeNames) {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) AcceptingCaseInsensitiveLike(typeName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	lower := strings.ToLower(typeName)

	for _, fa := range level {
		for _, paramType := range fa.TypesParams {
			if strings.Contains(strings.ToLower(paramType), lower) {
				result = append(result, fa)

				break // only one entry per function if multiple params match.
			}
		}
	}

	return result
}

func (level LevelFunction) Returning(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if unorderedButContainsAll(fa.TypesResults, typeNames) {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) ReturningCaseInsensitiveLike(typeName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	lower := strings.ToLower(typeName)

	for _, fa := range level {
		for _, paramType := range fa.TypesResults {
			if strings.Contains(strings.ToLower(paramType), lower) {
				result = append(result, fa)

				break // only one entry per function if multiple params match.
			}
		}
	}

	return result
}

func (level LevelFunction) HasVariadic() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if fa.HasVariadic {
			result = append(result, fa)
		}
	}

	return result
}
