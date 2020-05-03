package scenario

import (
	"errors"
	iface "github.com/genome-engine/genome/engine/scenario/interfaces"
)

//A scenario is a strict sequence of parsing,
//execute intermediate commands and generate code.
type Scenario struct {
	Chain []Step
}

func New(chain []Step, analyzer iface.Analyzer) (*Scenario, error) {
	if analyzer.Analyze() != nil {
		return nil, analyzer.Analyze()
	}

	return &Scenario{chain}, nil
}

func (s *Scenario) Add(name StepName, exec iface.Executable) error {
	u := Unknown

	if name.String() != u.String() {
		s.Chain = append(s.Chain, Step{name, exec})
		return nil
	}

	return errors.New("Unknown step name: " + name.String())
}
