package interfaces

//The Analyzer interface is used to analyze the scenario correctness.
//The Analyze method will always be executed when creating a script,
//i.e. when scenario.New() is created.
type Analyzer interface {
	Analyze() error
}
