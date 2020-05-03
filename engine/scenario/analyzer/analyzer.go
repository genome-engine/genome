package analyzer

import (
	"errors"
	s "github.com/genome-engine/genome/engine/scenario"
	"github.com/genome-engine/genome/helpers/errors_merger"
	"strings"
)

const (
	sequenceCheck string = "SEQUENCE_CHECK_PHASE"
)

//This is a simple script analyzer that goes by default.
//Analyzes how the steps in a script are followed.
//You can use yours by implementing the appropriate interface.
type DefaultScenarioAnalyzer struct {
	Steps []s.Step
}

func New(steps []s.Step) *DefaultScenarioAnalyzer {
	return &DefaultScenarioAnalyzer{Steps: steps}
}

//In the default implementation, the Analyze method runs private methods of checking
//and combines the errors returned from them into one.
func (a *DefaultScenarioAnalyzer) Analyze() error {
	var errs = map[string]error{}
	var err error

	err = a.sequenceCheck()

	if err != nil {
		errs[sequenceCheck] = err
	}

	errsResult := errors_merger.MergeMapErrors(errs)

	if errsResult != "" {
		return errors.New(errsResult)
	}

	return nil
}

//checks for compliance with step sequence rules prescribed in the defaultSequenceOfSteps method
func (a *DefaultScenarioAnalyzer) sequenceCheck() error {
	var last s.StepName
	var errorText strings.Builder

	errorText.WriteString("Scenario step sequence error:\n\t")

	for i, current := range a.Steps {
		if i > 0 {
			if stepExist(last, current.StepName) {
				last = current.StepName
				continue
			}
			errorText.WriteString("step ")
			errorText.WriteString(current.String())
			errorText.WriteString(" can't follow after step ")
			errorText.WriteString(last.String())
			errorText.WriteString(". Possible steps:\n\t")
			for _, possible := range defaultSequenceOfSteps()[last] {
				errorText.WriteString(possible.String())
				errorText.WriteString(", ")
			}

			return errors.New(errorText.String())
		}
		last = current.StepName
	}

	return nil
}

//correct step sequence by default
func defaultSequenceOfSteps() map[s.StepName][]s.StepName {
	all := []s.StepName{s.Generating, s.Parsing, s.Command}

	return map[s.StepName][]s.StepName{
		s.Unknown:    {},
		s.Parsing:    all,
		s.Command:    {s.Parsing, s.Command},
		s.Generating: all,
	}
}

func stepExist(current, needle s.StepName) bool {
	for _, step := range defaultSequenceOfSteps()[current] {
		if step == needle {
			return true
		}
	}
	return false
}
