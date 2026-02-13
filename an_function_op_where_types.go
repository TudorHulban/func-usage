package funcusage

import "slices"

func (a LevelFunction) WithNoParams() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if len(usage.TypesParams) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WithNoResults() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if len(usage.TypesResults) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) WithErrorReturn() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if slices.Contains(usage.TypesResults, "error") {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) AcceptingOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if unorderedButSameItems(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) ReturningOnly(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if unorderedButSameItems(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) Accepting(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if unorderedButContainsAll(usage.TypesParams, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) Returning(typeNames ...string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if unorderedButContainsAll(usage.TypesResults, typeNames) {
			result = append(result, usage)
		}
	}

	return result
}
