package funcusage

// AnalyzeMode defines how test files influence the analysis.
// Only one mode is active at a time. Each mode represents a distinct
// perspective on how production code and test code interact.
//
// The modes are intentionally *mutually exclusive* to avoid boolean-flag
// combinatorics and to keep the API deterministic and predictable.
type AnalyzeMode int

const (
	// ModeDefault analyzes only production code.
	// Test files are ignored entirely:
	//   - test functions are not reported
	//   - calls from tests do not increase usage counts
	//
	// This mode is ideal for pure dead-code detection in production code.
	ModeDefault AnalyzeMode = iota + 1

	// ModeIncludeTestsForCoverage counts calls from test files *into*
	// production code, but does not report test-defined functions.
	//
	// This mode is useful when you want a coverage-like view of which
	// production functions are exercised by tests, without treating test
	// helpers as part of the API surface.
	ModeIncludeTestsForCoverage

	// ModeIncludeTestHelpers includes functions defined in test files
	// (e.g. *_test.go) in the output, and counts their usage normally.
	//
	// This mode is useful for cleaning up test suites, identifying unused
	// test helpers, and understanding the internal API surface of tests.
	ModeIncludeTestHelpers

	// ModeOnlyTestHelpers analyzes *only* functions defined in test files.
	// Production code is ignored entirely.
	//
	// This mode is ideal for:
	//   - finding unused test helpers
	//   - understanding test-only utility layers
	//   - refactoring large test suites
	ModeOnlyTestHelpers

	// ModeOnlyInTestFiles counts only calls that *originate* from test files.
	// Production functions are still reported, but only with test-origin
	// internal/external counts.
	//
	// This mode is useful for:
	//   - identifying production functions used exclusively by tests
	//   - detecting over-exported APIs that are not used by real consumers
	//   - understanding test-to-prod dependency patterns
	ModeOnlyInTestFiles
)
