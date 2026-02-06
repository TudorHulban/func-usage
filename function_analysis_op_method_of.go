package funcusage

import (
	"strings"
)

func (a Analysis) IsFunction() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if len(usage.MethodOf) == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) IsMethod() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if len(usage.MethodOf) > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) MethodOf(objectName string) Analysis {
	result := make(Analysis, 0, len(a))

	// Trim leading * for comparison
	objectName = strings.TrimPrefix(objectName, "*")

	for _, usage := range a {
		if strings.TrimPrefix(usage.MethodOf, "*") == objectName {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) MethodLike(substr string) Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if strings.Contains(usage.MethodOf, substr) {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) MethodOfPointerReceiver() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if strings.HasPrefix(usage.MethodOf, "*") {
			result = append(result, usage)
		}
	}

	return result
}

func (a Analysis) MethodOfValueReceiver() Analysis {
	result := make(Analysis, 0, len(a))

	for _, usage := range a {
		if usage.MethodOf != "" && !strings.HasPrefix(usage.MethodOf, "*") {
			result = append(result, usage)
		}
	}

	return result
}
