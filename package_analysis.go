package funcusage

type AnalysisPackage struct {
	Name  string
	Types map[string]bool
}

type analysis struct {
	LevelFunction []AnalysisFunction
	LevelPackage  []AnalysisPackage
}
