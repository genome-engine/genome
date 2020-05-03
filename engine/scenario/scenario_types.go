package scenario

type Type int

const (
	Unknown Type = iota
	Command
	Parsing
	Generating
)

func (t *Type) String() string {
	mapping := map[Type]string{
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
