package funcusage

func (a LevelFunction) Where(predicate func(AnalysisFunction) bool) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if predicate(usage) {
			result = append(result, usage)
		}
	}

	return result
}
