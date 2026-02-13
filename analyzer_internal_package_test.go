package funcusage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackage(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	require.NotEmpty(t, analysis.LevelPackage)

	fmt.Println(
		analysis.LevelPackage,
	)
}
