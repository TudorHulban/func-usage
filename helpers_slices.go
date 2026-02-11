package funcusage

func unorderedButSameItems[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int, len(a))

	for _, valueA := range a {
		counts[valueA]++
	}

	for _, valueB := range b {
		if counts[valueB] == 0 {
			return false
		}

		counts[valueB]--
	}

	return true
}

func unorderedButContainsAll[T comparable](have, want []T) bool {
	if len(want) == 0 {
		return true
	}

	counts := make(map[T]int, len(have))

	for _, valueHave := range have {
		counts[valueHave]++
	}

	for _, valueWant := range want {
		if counts[valueWant] == 0 {
			return false
		}

		counts[valueWant]--
	}

	return true
}
