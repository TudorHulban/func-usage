package funcusage

import "slices"

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
