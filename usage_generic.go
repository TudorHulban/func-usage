package funcusage

func (u Usage) Where(predicate func(FunctionUsage) bool) Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if predicate(usage) {
			result = append(result, usage)
		}
	}

	return result
}
