package temp_env

import (
	"github.com/genome-engine/genome/engine/units"
)

//The wrappers above the structures of the units package.
type (
	Filterer struct {
		k string
		v []interface{}
	}
	ImportList    []units.Import
	FunctionList  []units.Function
	StructureList []units.Structure
	VariableList  []units.Variable
	InterfaceList []units.Interface

	UnknownList []units.Unknown
	MethodList  []units.Method

	PackageList  []units.Package
	CustomList   []units.Custom
	ConstantList []units.Constant
)

func NewFilterer(k string, v ...interface{}) Filterer { return Filterer{k: k, v: v} }
