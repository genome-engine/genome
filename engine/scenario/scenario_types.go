package scenario

import iface "github.com/genome-engine/genome/engine/scenario/interfaces"

type Step struct {
	StepName
	iface.Executable
}

//type is used to name the scenario steps
type StepName int

//possible steps
const (
	Unknown StepName = iota
	Command
	Parsing
	Generating
)

//if not on the list, it will return "Unknown"
func (t *StepName) String() string {
	mapping := map[StepName]string{
		Unknown:    "Unknown",
		Command:    "Command",
		Parsing:    "Parsing",
		Generating: "Generating",
	}

	name, ok := mapping[*t]

	if ok {
		return name
	}
	return mapping[Unknown]
}
