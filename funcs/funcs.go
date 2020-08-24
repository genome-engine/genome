package funcs

import (
	"github.com/genome-engine/genome/temp_env"
	"strings"
)

var Funcs = map[string]interface{}{
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
