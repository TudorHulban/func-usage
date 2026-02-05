package funcusage

// built from primitives

func (u Usage) MostUsed(n int) Usage {
	return u.OrderByTotalCallsDesc().Limit(n)
}

func (u Usage) LeastUsed(n int) Usage {
	return u.OrderByTotalCallsAsc().Limit(n)
}

func (u Usage) ExportedWithNoExternalCalls() Usage {
	return u.WhereExported().Where(func(fn FunctionUsage) bool {
		return fn.ExternalCount == 0 && fn.InternalCount > 0
	})
}

func (u Usage) ExportedUnused() Usage {
	return u.WhereExported().WhereUnused()
}
