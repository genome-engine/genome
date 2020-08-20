package script

type Delimiter int

const (
	Arrow Delimiter = iota
	Braces
)

var delimiters = map[Delimiter][]string{Arrow: {"<", ">"}, Braces: {"{{", "}}"}}

func (d Delimiter) Delimiters() (string, string) { return delimiters[d][0], delimiters[d][1] }
