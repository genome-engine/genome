package scenario

import (
	"errors"
	iface "github.com/genome-engine/genome/engine/scenario/interfaces"
	"strings"
)

//A scenario is a strict sequence of parsing,
//execute intermediate commands and generate code.
type Scenario struct {
	Chain []Step
	iface.Analyzer
}

func New(chain []Step, analyzer iface.Analyzer) (*Scenario, error) {
	return &Scenario{chain, analyzer}, nil
}

func (s *Scenario) Add(name StepName, exec iface.Executable) error {
	u := Unknown

	if name.String() != u.String() {
		s.Chain = append(s.Chain, Step{name, exec})
		return nil
	}

	return errors.New("Unknown step name: " + name.String())
}

func (s *Scenario) Execute() error {
	var errText strings.Builder

	err := s.Analyze()

	if err != nil {
		return err
	}

	for _, executable := range s.Chain {
		err = executable.Execute()
		if err != nil {
			errText.WriteString(strings.ToUpper(executable.String()))
			errText.WriteString("_PHASE_ERROR:\n\t")
			errText.WriteString(err.Error())
		}
	}

	errorsResult := errText.String()

	if errorsResult != "" {
		return errors.New(errorsResult)
	}

	return nil
}
