package funcusage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnalysis(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	usage, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, usage)

	require.NotEmpty(t,
		usage.IsMethod(),
	)

	printer := NewPrinter().WithName().WithMethodOf()

	usage.
		MethodOf("Analysis").
		OrderByNameAsc().
		PrintWith(printer)
}

func TestUsageFilters(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	usage, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
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
			require.NotEmpty(t,
				usage.WhereExported(),
			)
		},
	)

	t.Run(
		"3. Limit - test zero limit",
		func(t *testing.T) {
			require.Empty(t,
				usage.WhereNameIs(fnName).Limit(0),
			)
		},
	)

	t.Run(
		"4. Limit - test different than zero value",
		func(t *testing.T) {
			require.Len(t,
				usage.WhereNameIs(fnName).Limit(1),
				1,
			)
		},
	)

	t.Run(
		"5. Limit - test chaining",
		func(t *testing.T) {
			require.Len(t,
				usage.
					WhereTestedInternally().
					WhereNameIs(fnName).
					Limit(1),
				1,
			)
		},
	)

	t.Run(
		"6. Limit - beyond length",
		func(t *testing.T) {
			require.Len(t,
				usage.WhereNameIs(fnName).Limit(100),
				len(usage.WhereNameIs(fnName)),
			)
		},
	)
}
