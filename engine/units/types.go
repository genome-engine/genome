package units

import (
	"github.com/genome-engine/genome/engine/types"
)

type (
	Description struct {
		selector   Selector
		Comment    string
		IsExported bool
		name       string
		id         int
		types.Type
	}
	Package struct {
		Description

		IsMain bool
		Path   string
	}
	Import struct {
		Description
		Value string
	}
	Structure struct {
		Description
		Fields []StructField
	}
	StructField struct {
		Comment    string
		Tag        string
		IsExported bool
		Name       string
		Type       types.Type
	}
	Interface struct {
		Description
		IsStructField bool
	}
	Custom struct {
		Description
	}
	Method struct {
		Description
		InInterfaceDecl     bool
		Signature           string
		Parameters, Returns map[string]types.Type
		Body                string
	}
	Function struct {
		Description
		Signature           string
		Parameters, Returns map[string]types.Type
		FuncBody            string
	}
	EnumSeries struct {
		Description
	}
	Constant struct {
		Description
		IsEnum bool
	}
	Variable struct {
		Description
	}
	Unknown struct {
		Description
	}
)
