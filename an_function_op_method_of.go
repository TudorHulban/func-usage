package funcusage

import (
	"strings"
)

func (level LevelFunction) IsFunction() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if len(usage.MethodOf) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) IsMethod() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if len(usage.MethodOf) > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) MethodOf(objectName string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	// Trim leading * for comparison
	objectName = strings.TrimPrefix(objectName, "*")

	for _, usage := range level {
		if strings.TrimPrefix(string(usage.MethodOf), "*") == objectName {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) MethodLike(substr string) LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if strings.Contains(string(usage.MethodOf), substr) {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) MethodOfPointerReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if strings.HasPrefix(string(usage.MethodOf), "*") {
			result = append(result, usage)
		}
	}

	return result
}

func (level LevelFunction) MethodOfValueReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(level))

	for _, usage := range level {
		if usage.MethodOf != "" && !strings.HasPrefix(string(usage.MethodOf), "*") {
			result = append(result, usage)
		}
	}

	return result
}
