package units

import (
	"github.com/genome-engine/genome/engine/types"
)

func NewPackage(id int, name string) *Package {
	return &Package{Description: Description{id: id, name: name, selector: GoPackage, Type: types.NewPackageType(name)}}
}
func NewImport(id int, name string) *Import {
	return &Import{Description: Description{id: id, name: name, selector: GoImport, Type: types.NewImportType(name)}}
}
func NewStruct(id int, name string) *Structure {
	return &Structure{Description: Description{id: id, name: name, selector: GoStruct, Type: types.NewStructType(name)}}
}
func NewIface(id int, name string) *Interface {
	return &Interface{Description: Description{id: id, name: name, selector: GoInterface, Type: types.NewIfaceType(name)}}
}
func NewCustom(id int, name string) *Custom {
	return &Custom{Description: Description{id: id, name: name, selector: GoCustom, Type: types.NewCustomType(name)}}
}
func NewMethod(id int, name string) *Method {
	return &Method{Description: Description{id: id, name: name, selector: GoMethod}}
}
func NewFunc(id int, name string) *Function {
	return &Function{Description: Description{id: id, name: name, selector: GoFunc}}
}
func NewEnumSeries() *EnumSeries {
	return &EnumSeries{Description: Description{selector: GoEnumSeries}}
}
func NewConst(id int, name string) *Constant {
	return &Constant{Description: Description{id: id, name: name, selector: GoConst}}
}
func NewVar(id int, name string) *Variable {
	return &Variable{Description: Description{id: id, name: name, selector: GoVar}}
}
func NewUnknown(id int, name string) *Unknown {
	return &Unknown{Description: Description{id: id, name: name, selector: GoUnknown, Type: types.NewUnknownType(name)}}
}

func Init(id int, name string, selector Selector) Unit {
	switch selector {
	case GoPackage:
		return NewPackage(id, name)
	case GoImport:
		return NewImport(id, name)
	case GoStruct:
		return NewStruct(id, name)
	case GoInterface:
		return NewIface(id, name)
	case GoCustom:
		return NewCustom(id, name)
	case GoMethod:
		return NewMethod(id, name)
	case GoFunc:
		return NewFunc(id, name)
	case GoConst:
		return NewConst(id, name)
	case GoVar:
		return NewVar(id, name)
	case GoUnknown:
		return NewUnknown(id, name)
	}

	return nil
}
