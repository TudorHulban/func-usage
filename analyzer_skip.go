package funcusage

func skipFileByMode(isTest bool, mode AnalyzeMode) bool {
	switch mode {
	case ModeDefault:
		return isTest

	case ModeOnlyTestHelpers, ModeOnlyInTestFiles:
		return !isTest

	default:
		return false
	}
}
