package funcusage

import "sort"

func (a Analysis) OrderByTotalCallsDesc() Analysis {
	result := make(Analysis, len(a))
	copy(result, a)

	sort.Slice(
		result, func(i, j int) bool {
			ti := result[i].InternalCount + result[i].ExternalCount
			tj := result[j].InternalCount + result[j].ExternalCount

			return ti > tj
		},
	)

	return result
}

func (a Analysis) OrderByTotalCallsAsc() Analysis {
	result := make(Analysis, len(a))
	copy(result, a)

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

func (a Analysis) OrderByExternalCallsDesc() Analysis {
	result := make(Analysis, len(a))
	copy(result, a)

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

func (a Analysis) OrderByNameAsc() Analysis {
	result := make(Analysis, len(a))
	copy(result, a)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name < result[j].Name
		},
	)

	return result
}

func (a Analysis) OrderByNameDesc() Analysis {
	result := make(Analysis, len(a))
	copy(result, a)

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Name > result[j].Name
		},
	)

	return result
}
