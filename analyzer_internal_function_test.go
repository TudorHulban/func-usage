package funcusage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLimit(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	fnSought := "Limit"

	assert.NotEmpty(t,
		analysis.
			LevelFunction.
			IsMethod().
			WhereNameIs(fnSought),
	)

	require.Empty(t,
		analysis.
			LevelFunction.
			IsFunction().
			WhereNameIs(fnSought),
	)

	fmt.Println(
		analysis.
			LevelFunction.
			IsMethod().
			WhereNameIs(fnSought).String(),
	)
}

func TestNoGroupingAnalysis(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	require.NotEmpty(t,
		analysis.LevelFunction.IsMethod(),
	)

	printer := NewPrinter().
		WithName().
		WithMethodOf().
		WithTypesParams().
		WithTypesResults()

	analysis.
		LevelFunction.
		MethodOf("Analysis").
		OrderByNameAsc().
		PrintWith(printer)
}

func TestWithGroupingAnalysis(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	require.NotEmpty(t,
		analysis.LevelFunction.IsMethod(),
	)

	printer := NewPrinter().WithName()

	analysis.
		LevelFunction.
		GroupedByObject().
		PrintWith(printer)
}

func TestDoubleGroupingAnalysis(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	require.NotEmpty(t,
		analysis.
			LevelFunction.
			IsMethod(),
	)

	printer := NewPrinter().WithName()

	analysis.
		LevelFunction.
		GroupedByPackageAndObject().
		PrintWith(printer)
}

func TestUsageFilters(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	fnName := "Analyze"

	t.Run(
		"1. WhereNameIs finds Analyze",
		func(t *testing.T) {
			require.NotEmpty(t,
				analysis.
					LevelFunction.
					WhereNameIs(fnName),
			)
		},
	)

	t.Run(
		"2. WhereExported finds exports",
		func(t *testing.T) {
			require.NotEmpty(t,
				analysis.
					LevelFunction.
					WhereExported(),
			)
		},
	)

	t.Run(
		"3. Limit - test zero limit",
		func(t *testing.T) {
			require.Empty(t,
				analysis.
					LevelFunction.
					WhereNameIs(fnName).
					Limit(0),
			)
		},
	)

	t.Run(
		"4. Limit - test different than zero value",
		func(t *testing.T) {
			require.Len(t,
				analysis.
					LevelFunction.
					WhereNameIs(fnName).
					Limit(1),
				1,
			)
		},
	)

	t.Run(
		"5. Limit - test chaining",
		func(t *testing.T) {
			require.Len(t,
				analysis.
					LevelFunction.
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
				analysis.
					LevelFunction.
					WhereNameIs(fnName).
					Limit(100),

				len(analysis.
					LevelFunction.
					WhereNameIs(fnName),
				),
			)
		},
	)
}

func TestSignatureGrouping(t *testing.T) {
	a, errCr := NewAnalyzer(".")
	require.NoError(t, errCr)
	require.NotNil(t, a)

	analysis, errAnalyze := a.Analyze(
		ModeIncludeTestsForCoverage,
		false,
	)
	require.NoError(t, errAnalyze)
	require.NotZero(t, analysis)

	require.NotEmpty(t,
		analysis.
			LevelFunction.
			IsMethod(),
	)

	printer := NewPrinter().
		WithName().
		WithMethodOf().
		WithTypesParams().
		WithTypesResults()

	analysis.
		LevelFunction.
		GroupedByParamSignature().
		PrintWith(printer)
}
