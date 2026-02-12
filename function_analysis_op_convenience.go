package funcusage

// built from primitives

func (a Analysis) MostUsed(n int) Analysis {
	return a.OrderByTotalCallsDesc().Limit(n)
}

func (a Analysis) LeastUsed(n int) Analysis {
	return a.OrderByTotalCallsAsc().Limit(n)
}

func (a Analysis) ExportedWithNoExternalCalls() Analysis {
	return a.WhereExported().Where(func(fn AnalysisFunction) bool {
		return fn.ExternalCount == 0 && fn.InternalCount > 0
	})
}

func (a Analysis) ExportedUnused() Analysis {
	return a.WhereExported().WhereUnused()
}
