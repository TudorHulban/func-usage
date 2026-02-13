package funcusage

// built from primitives

func (level LevelFunction) MostUsed(n int) LevelFunction {
	return level.OrderByTotalCallsDesc().Limit(n)
}

func (level LevelFunction) LeastUsed(n int) LevelFunction {
	return level.OrderByTotalCallsAsc().Limit(n)
}

func (level LevelFunction) ExportedWithNoExternalCalls() LevelFunction {
	return level.WhereExported().Where(func(fn AnalysisFunction) bool {
		return fn.ExternalCount == 0 && fn.InternalCount > 0
	})
}

func (level LevelFunction) ExportedUnused() LevelFunction {
	return level.WhereExported().WhereNotUsed()
}
