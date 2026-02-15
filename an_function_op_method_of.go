package funcusage

import (
	"strings"
)

func (level LevelFunction) IsFunction() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if len(fa.MethodOf) == 0 {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) IsMethod() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if len(fa.MethodOf) > 0 {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) MethodOf(objectName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	// Trim leading * for comparison
	objectName = strings.TrimPrefix(objectName, "*")

	for _, fa := range level {
		if strings.TrimPrefix(string(fa.MethodOf), "*") == objectName {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) MethodLike(substr string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if strings.Contains(string(fa.MethodOf), substr) {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) MethodOfPointerReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if strings.HasPrefix(string(fa.MethodOf), "*") {
			result = append(result, fa)
		}
	}

	return result
}

func (level LevelFunction) MethodOfValueReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, fa := range level {
		if fa.MethodOf != "" && !strings.HasPrefix(string(fa.MethodOf), "*") {
			result = append(result, fa)
		}
	}

	return result
}
