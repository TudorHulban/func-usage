package funcusage

import (
	"slices"
	"strings"
)

func (level LevelFunction) WithNoParams() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if len(usage.TypesParams) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WithNoResults() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if len(usage.TypesResults) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) WithErrorReturn() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if slices.Contains(usage.TypesResults, "error") {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) AcceptingOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if unorderedButSameItems(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) ReturningOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if unorderedButSameItems(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) Accepting(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if unorderedButContainsAll(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) AcceptingCaseInsensitiveLike(typeName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	lower := strings.ToLower(typeName)

	for _, usage := range level {
		for _, paramType := range usage.TypesParams {
			if strings.Contains(strings.ToLower(paramType), lower) {
				result = append(result, usage)

				break // only one entry per function if multiple params match.
			}
		}
	}

	return result
}

func (level LevelFunction) Returning(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if unorderedButContainsAll(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) ReturningCaseInsensitiveLike(typeName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	lower := strings.ToLower(typeName)

	for _, usage := range level {
		for _, paramType := range usage.TypesResults {
			if strings.Contains(strings.ToLower(paramType), lower) {
				result = append(result, usage)

				break // only one entry per function if multiple params match.
			}
		}
	}

	return result
}
