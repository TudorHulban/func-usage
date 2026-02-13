package funcusage

// built from primitives

func (a LevelFunction) MostUsed(n int) LevelFunction {
	return a.OrderByTotalCallsDesc().Limit(n)
}

func (a LevelFunction) LeastUsed(n int) LevelFunction {
	return a.OrderByTotalCallsAsc().Limit(n)
}

func (a LevelFunction) ExportedWithNoExternalCalls() LevelFunction {
	return a.WhereExported().Where(func(fn AnalysisFunction) bool {
		return fn.ExternalCount == 0 && fn.InternalCount > 0
	})
}

func (a LevelFunction) ExportedUnused() LevelFunction {
	return a.WhereExported().WhereUnused()
}
