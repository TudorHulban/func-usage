package funcusage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVariadic(t *testing.T) {
	analyzer, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, analyzer)

	analysis, errAnalyze := analyzer.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	fa := analysis.LevelFunction.HasVariadic()

	require.NotEmpty(t, fa)
	require.True(t,
		fa.ContainsAll(
			"AcceptingOnly",
		),
	)
}
