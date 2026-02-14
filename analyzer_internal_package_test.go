package funcusage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatistics(t *testing.T) {
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
		"number functions:",
		len(analysis.LevelFunction.IsFunction()),
	)

	fmt.Println(
		"number methods:",
		len(analysis.LevelFunction.IsMethod()),
	)

	fmt.Println(
		"number untested:",
		len(analysis.LevelFunction.WhereNotTested()),
	)

	fmt.Println(
		"number unused:",
		len(analysis.LevelFunction.WhereNotUsed()),
	)

	fmt.Println(
		analysis.
			LevelPackage.
			Statistics(a.ModulePath),
	)
}
