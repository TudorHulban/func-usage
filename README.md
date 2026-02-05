# Function Usage

Function Usage is **not** a standalone CLI tool — it is a lightweight, composable static analysis library designed to map how functions are used across a Go repository and live inside tests.  
It provides a clear, type‑checked view of call relationships, allowing one to understand usage patterns and identify structural issues in your codebase.  
Function Usage is intentionally simple, deterministic, and transparent.
It relies on Go’s type checker rather than heuristics or heavy call‑graph machinery, giving one predictable and actionable results.  

It can help with below:

a. find most used functions  
b. find least used functions  
c. find exported functions not used  
d. find exported functions that could be unexported  
e. have the posibility to include or exclude tests or external functions.  

## How to use

The test files should provide examples on how to use.  
A possible workflow could be:

### Analyze the repository

The `Usage` data returned by the analyzer represents an expensive snapshot of the codebase (AST parsing, type checking). All filter and sort operations preserve this original data by returning new slices.  
Run the analyzer from the module root as a test.  
This way it can run both in CI and also resolve the module path and load all packages, including tests.

#### Performance Trade-off

- **Analysis Phase**: Expensive (seconds, type checking)
- **Query Phase**: Cheap (microseconds, slice operations)
- **Design Choice**: Accept O(n) copying in where operations to guarantee:
  - Test determinism (no flaky tests from mutations)
  - Safe method chaining
  - Predictable debugging (original data always available)

#### Scope Control

Configure whether to include:

- same‑package tests
- external test packages (`mypkg_test`)
- external packages outside the module

#### Usage Structure

The analyzer returns a `Usage` slice.  
Each `FunctionUsage` entry contains:

- canonical function identity
- short name
- source position
- internal and external call counts
- internal and external test call counts

### Chaining Filters

Filters return a new `Usage`, allowing composition:

```go
untestedExported := usage.
    WhereNotTested().
    WhereExported().
    Limit(10)
```

Example: find a specific untested exported function:

```go
result := usage.
    WhereNotTested().
    WhereExported().
    WhereNameIs("some function name")
```

## Modes

`funcusage` provides several mutually exclusive analysis modes. Each mode defines how production code and test code are interpreted during analysis.

### ModeDefault

Analyze only production code.

- Test files are ignored.
- Test-defined functions are not reported.
- Calls from tests do not count toward usage.

Use this mode for strict dead‑code detection in production.

### ModeIncludeTestsForCoverage

Include calls from test files into production code, but do not report test-defined functions.

- Test helpers are ignored.
- Calls from tests increase usage counts for production functions.

Use this mode for a coverage‑like view of how tests exercise production code.

### ModeIncludeTestHelpers

Include both production functions and test-defined functions.

- Test helpers appear in the output.
- Their usage is counted normally.

Use this mode to analyze or clean up test suites.

### ModeOnlyTestHelpers

Analyze only functions defined in test files.

- Production code is ignored.
- Only test helpers and their usage are reported.

Use this mode to audit or refactor large test suites.

### ModeOnlyInTestFiles

Report production functions, but count only calls originating from test files.

- Production functions are included.
- Only test-origin usage counts are considered.

Use this mode to identify production functions used exclusively by tests or over‑exported APIs.

## Filters

Filters allow narrowing the analysis result (`Usage`) based on name, visibility, or usage patterns.  
The number of values can be controlled with `Limit` as last element in the chain.  
All filter methods start with `Where` and return new `Usage` slices:

### WhereNameIs

Return only functions whose name matches exactly passed name.

### WhereUnused

Return functions with zero internal and zero external calls.

### WhereExported

Return functions whose name starts with an upper‑case letter.

### WhereUnexported

Return functions whose name starts with a lower‑case letter.

### WhereTestedInternally

Return functions that have internal test calls.  

### WhereTestedExternally

Return functions that have external test calls.  

### ExportedUnused

Return exported functions that have no internal or external calls.

### ExportedWithNoExternalCalls

Return exported functions that have no external calls.  
Useful for identifying candidates that should be unexported.

### Generic Where

Allows injecting custom predicate for extensibility.

## Sorters / Order by

### OrderByTotalCallsDesc

Sort by total calls (`internal + external`), highest first.  
If totals are equal, sort by key.

### OrderByTotalCallsAsc

Sort by number of calls, lowest number of calls first.  
If counts are equal, sort by key.

### OrderByExternalCallsDesc

Sort by external calls, highest number of calls first.  

### OrderByNameAsc

Sort by function name A → Z.

### OrderByNameDesc

Sort by function name Z → A.

### MostUsed

Returns the first N highest number of calls functions.

### LeastUsed

Returns the first N lowest number of calls functions.

## Declared Functions Scope (Intentional Limitation)

funcusage analyzes declared functions and methods (func declarations) only.

It intentionally does not track:

- function literals or closures
- anonymous functions assigned to variables
- dynamically constructed call targets
- method usage through interfaces.

This is a deliberate design choice.  
Declared functions form the stable, addressable API surface of a Go codebase — the part that is exported, refactored, reviewed, and reasoned about at scale.  
Including function literals would significantly increase noise while providing little actionable insight for structural analysis.

As a result, funcusage may undercount usage in highly functional or closure-heavy code, but it avoids false positives and preserves deterministic, explainable results.
