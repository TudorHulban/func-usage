package funcusage

func (level LevelFunction) Where(predicate func(AnalysisFunction) bool) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if predicate(*fa) {
			result = append(result, fa)
		}
	}

	return result
}
