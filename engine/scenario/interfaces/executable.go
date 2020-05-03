package interfaces

//The Executable interface is used to combine the execution of steps in a script.
//That is for Command, Generating, Parsing (you can come up with your own).
type Executable interface {
	Execute() error
}
