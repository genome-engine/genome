package temp_env

import (
	"github.com/genome-engine/genome/engine/units"
)

//The wrappers above the structures of the units package.
type (
	ImportList    []units.Import
	FunctionList  []units.Function
	StructureList []units.Structure
	VariableList  []units.Variable
	InterfaceList []units.Interface
	UnknownList   []units.Unknown
	MethodList    []units.Method
	PackageList   []units.Package
	CustomList    []units.Custom
	ConstantList  []units.Constant
)

//#genome-insert:list
func (list *FunctionList) List() []units.Function   { return *list }
func (list *VariableList) List() []units.Variable   { return *list }
func (list *InterfaceList) List() []units.Interface { return *list }
func (list *StructureList) List() []units.Structure { return *list }
func (list *UnknownList) List() []units.Unknown     { return *list }
func (list *MethodList) List() []units.Method       { return *list }
func (list *PackageList) List() []units.Package     { return *list }
func (list *ConstantList) List() []units.Constant   { return *list }
func (list *CustomList) List() []units.Custom       { return *list }

//#genome-insert-end
