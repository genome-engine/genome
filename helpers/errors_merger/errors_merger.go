package errors_merger

import (
	"strconv"
	"strings"
)

//aggregates the error text into a single string
func MergeMapErrors(errs map[string]error) string {
	var errText strings.Builder

	for phase, err := range errs {
		errText.WriteString(phase)
		errText.WriteString(":\n\t")
		errText.WriteString(err.Error())
	}

	return errText.String()
}

func MergeArrayErrors(errs []error) string {
	var errText strings.Builder

	for i, err := range errs {
		errText.WriteString(strconv.FormatInt(int64(i), 10))
		errText.WriteString(".")
		errText.WriteString(err.Error())
	}

	return errText.String()
}
