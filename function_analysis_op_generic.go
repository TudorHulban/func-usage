package funcusage

func (a Analysis) Where(predicate func(FunctionAnalysis) bool) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if predicate(usage) {
			result = append(result, usage)
		}
	}

	return result
}
