package funcusage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	funcusage "github.com/tudorhulban/func-usage"
)

func TestExternalAnalyzer(t *testing.T) {
	a, errCr := funcusage.NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		funcusage.ModeIncludeTestsForCoverage,
		true,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	fnName := "Analyze"
	packageName := "funcusage"

	t.Run(
		"1. WhereNameIs finds Analyze",
		func(t *testing.T) {
			require.NotEmpty(t,
				analysis.LevelFunction.WhereNameIs(fnName),
			)
		},
	)

	t.Run(
		"2. WhereExported finds exports",
		func(t *testing.T) {
			require.Empty(t,
				analysis.
					LevelFunction.
					WhereNameIs(fnName).
					WhereUnexported(),
			)
		},
	)

	t.Run(
		"3. Limit - test chaining",
		func(t *testing.T) {
			require.Len(t,
				analysis.
					LevelFunction.
					WhereTestedExternally().
					WhereNameIs(fnName).
					Limit(1),
				1,
			)
		},
	)

	t.Run(
		"4. WherePackageIs",
		func(t *testing.T) {
			require.Len(t,
				analysis.
					LevelFunction.
					WhereTestedExternally().
					WherePackageIs(packageName).
					WhereNameIs(fnName).
					Limit(1),
				1,
			)
		},
	)
}
