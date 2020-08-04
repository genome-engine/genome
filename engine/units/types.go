package units

import (
	"github.com/genome-engine/genome/engine/types"
)

type (
	Package struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		Main bool
		Path string
	}
	Import struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		Value string
	}
	Structure struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
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
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		IsStructField bool
	}
	Custom struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
	}
	Method struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		InInterfaceDecl     bool
		Signature           string
		Parameters, Returns map[string]types.Type
		Body                string
	}
	Function struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		Signature           string
		Parameters, Returns map[string]types.Type
		FuncBody            string
	}
	Constant struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
		Enum bool
	}
	Variable struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
	}
	Unknown struct {
		Selector   Selector
		Comment    string
		IsExported bool
		Name       string
		ID         int
		types.Type
	}
)
