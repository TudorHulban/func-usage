package funcusage

import "sort"

func (level LevelFunction) OrderByTotalCallsDesc() LevelFunction {
	result := make(LevelFunction, len(level))
	copy(result, level)

	sort.Slice(
		result, func(i, j int) bool {
			ti := result[i].InternalCount + result[i].ExternalCount
			tj := result[j].InternalCount + result[j].ExternalCount

			return ti > tj
		},
	)

	return result
}

func (level LevelFunction) OrderByTotalCallsAsc() LevelFunction {
	result := make(LevelFunction, len(level))
	copy(result, level)

	sort.Slice(
		result,
		func(i, j int) bool {
			ti := result[i].InternalCount + result[i].ExternalCount
			tj := result[j].InternalCount + result[j].ExternalCount

			return ti < tj
		},
	)

	return result
}

func (level LevelFunction) OrderByExternalCallsDesc() LevelFunction {
	result := make(LevelFunction, len(level))
	copy(result, level)

	sort.Slice(
		result,
		func(i, j int) bool {
			if result[i].ExternalCount == result[j].ExternalCount {
				return result[i].Key < result[j].Key
			}

			return result[i].ExternalCount > result[j].ExternalCount
		},
	)

	return result
}

func (level LevelFunction) OrderByNameAsc() LevelFunction {
	result := make(LevelFunction, len(level))
	copy(result, level)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}

func (level LevelFunction) OrderByNameDesc() LevelFunction {
	result := make(LevelFunction, len(level))
	copy(result, level)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name > result[j].Name
		},
	)

	return result
}
