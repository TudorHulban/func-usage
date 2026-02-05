package funcusage

import "sort"

func (u Usage) OrderByTotalCallsDesc() Usage {
	result := make(Usage, len(u))
	copy(result, u)

	sort.Slice(
		result, func(i, j int) bool {
			ti := result[i].InternalCount + result[i].ExternalCount
			tj := result[j].InternalCount + result[j].ExternalCount

			return ti > tj
		},
	)

	return result
}

func (u Usage) OrderByTotalCallsAsc() Usage {
	result := make(Usage, len(u))
	copy(result, u)

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

func (u Usage) OrderByExternalCallsDesc() Usage {
	result := make(Usage, len(u))
	copy(result, u)

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

func (u Usage) OrderByNameAsc() Usage {
	result := make(Usage, len(u))
	copy(result, u)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}

func (u Usage) OrderByNameDesc() Usage {
	result := make(Usage, len(u))
	copy(result, u)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name > result[j].Name
		},
	)

	return result
}
