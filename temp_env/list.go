package temp_env

import (
	"github.com/genome-engine/genome/engine/units"
)

//The wrappers above the structures of the units package.
type (
	UnknownList []units.Unknown
	ImportList  []units.Import
	StructList  []units.Struct
	MethodList  []units.Method
	CustomList  []units.Custom
	IfaceList   []units.Iface
	ConstList   []units.Const
	PackList    []units.Pack
	FuncList    []units.Func
	VarList     []units.Var
)

func (list *UnknownList) List() []units.Unknown { return *list }
func (list *ImportList) List() []units.Import   { return *list }
func (list *StructList) List() []units.Struct   { return *list }
func (list *MethodList) List() []units.Method   { return *list }
func (list *CustomList) List() []units.Custom   { return *list }
func (list *IfaceList) List() []units.Iface     { return *list }
func (list *ConstList) List() []units.Const     { return *list }
func (list *PackList) List() []units.Pack       { return *list }
func (list *FuncList) List() []units.Func       { return *list }
func (list *VarList) List() []units.Var         { return *list }
