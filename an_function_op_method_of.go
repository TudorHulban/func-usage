package funcusage

import (
	"strings"
)

func (a LevelFunction) IsFunction() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if len(usage.MethodOf) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) IsMethod() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if len(usage.MethodOf) > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) MethodOf(objectName string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	// Trim leading * for comparison
	objectName = strings.TrimPrefix(objectName, "*")

	for _, usage := range a {
		if strings.TrimPrefix(string(usage.MethodOf), "*") == objectName {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) MethodLike(substr string) LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if strings.Contains(string(usage.MethodOf), substr) {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) MethodOfPointerReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if strings.HasPrefix(string(usage.MethodOf), "*") {
			result = append(result, usage)
		}
	}

	return result
}

func (a LevelFunction) MethodOfValueReceiver() LevelFunction {
	result := make(LevelFunction, 0, len(a))

	for _, usage := range a {
		if usage.MethodOf != "" && !strings.HasPrefix(string(usage.MethodOf), "*") {
			result = append(result, usage)
		}
	}

	return result
}
