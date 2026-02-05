package funcusage

// Limit returns at most n elements from the slice.
// Should be used as the last element in a chain of operations.
//
// Why returning a slice (not a copy) is safe:
// 1. All preceding operations (Where*, OrderBy*) return NEW slices
// 2. Limit operates on a slice that exists ONLY within this chain
// 3. No other code holds references to this intermediate slice
// 4. FunctionUsage values are immutable (no reference types)
//
// Example safe usage:
//
//	result := usage.
//	    WhereExported().          // ← new slice
//	    OrderByTotalCallsDesc().  // ← new slice
//	    Limit(10)                 // ← slice of the last new slice
//
// Memory efficient: O(1) slice operation, no allocation.
// Follows Go's slice semantics (like u[:n]).
func (u Usage) Limit(n int) Usage {
	if n <= 0 {
		return Usage{}
	}

	if n >= len(u) {
		return u
	}

	return u[:n]
}
