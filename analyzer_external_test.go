package funcusage_test

import (
	"testing"

	funcusage "github.com/TudorHulban/func-usage"
	"github.com/stretchr/testify/require"
)

func TestExternalAnalyzer(t *testing.T) {
	a, errCr := funcusage.NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	usage, errAnalyze := a.Analyze(
		funcusage.ModeIncludeTestsForCoverage,
		true,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, usage)

	fnName := "Analyze"

	t.Run(
		"1. WhereNameIs finds Analyze",
		func(t *testing.T) {
			require.NotEmpty(t,
				usage.WhereNameIs(fnName),
			)
		},
	)

	t.Run(
		"2. WhereExported finds exports",
		func(t *testing.T) {
			require.Empty(t,
				usage.
					WhereNameIs(fnName).
					WhereUnexported(),
			)
		},
	)

	t.Run(
		"3. Limit - test chaining",
		func(t *testing.T) {
			require.Len(t,
				usage.
					WhereTestedExternally().
					WhereNameIs(fnName).
					Limit(1),
				1,
			)
		},
	)
}
