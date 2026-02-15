package funcusage

func (level LevelFunction) ContainsAll(fnName ...string) bool {
	if len(fnName) == 0 {
		return true
	}

	// approach is O(n + m) vs direct walk of slice which is O(n*m).
	set := make(map[string]struct{}, len(level))

	for _, fa := range level {
		set[fa.Name] = struct{}{}
	}

	for _, name := range fnName {
		if _, exists := set[name]; !exists {
			return false
		}
	}

	return true
}
