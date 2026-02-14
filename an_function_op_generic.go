package funcusage

func (level LevelFunction) Where(predicate func(AnalysisFunction) bool) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if predicate(*usage) {
			result = append(result, usage)
		}
	}

	return result
}
