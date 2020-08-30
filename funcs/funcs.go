package funcs

import (
	"fmt"
	"github.com/genome-engine/genome/plugin"
	"github.com/genome-engine/genome/temp_env"
	"strings"
)

var funcs = map[string]interface{}{
	"join":     strings.Join,
	"title":    strings.Title,
	"lower":    strings.ToLower,
	"upper":    strings.ToUpper,
	"contains": strings.Contains,
	"trim_r":   strings.TrimRight,
	"trim_l":   strings.TrimLeft,
	"trim":     strings.Trim,
	"f":        temp_env.NewFilter,
}

func Funcs() map[string]interface{} {
	for name, fun := range plugin.Load() {
		if _, ok := funcs[name]; !ok {
			funcs[name] = fun
		} else {
			fmt.Printf("\t\tFunction with name[%v] exist in builtin function list of genome.", name)
		}
	}

	return funcs
}
