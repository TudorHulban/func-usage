package funcusage

func (u Usage) WhereNameIs(name string) Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if usage.Name == name {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereUnused() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if usage.InternalCount == 0 && usage.ExternalCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereExported() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereUnexported() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if !isExportedName(usage.Name) {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereNotTested() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if usage.InternalTestsCount == 0 && usage.ExternalTestsCount == 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereTestedInternally() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if usage.InternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}

func (u Usage) WhereTestedExternally() Usage {
	result := make(Usage, 0, len(u))

	for _, usage := range u {
		if usage.ExternalTestsCount > 0 {
			result = append(result, usage)
		}
	}

	return result
}
