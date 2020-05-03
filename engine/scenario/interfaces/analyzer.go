package interfaces

//The Analyzer interface is used to analyze the scenario correctness.
//The Analyze method will always be executed when you call scenario.Execute
type Analyzer interface {
	Analyze() error
}
