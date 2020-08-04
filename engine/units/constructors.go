package units

import (
	"github.com/genome-engine/genome/engine/types"
)

func NewPackage(id int, name string) *Package {
	return &Package{ID: id, Name: name, Selector: GoPackage, Type: types.NewPackageType(name)}
}
func NewImport(id int, name string) *Import {
	return &Import{ID: id, Name: name, Selector: GoImport, Type: types.NewImportType(name)}
}
func NewStruct(id int, name string) *Structure {
	return &Structure{ID: id, Name: name, Selector: GoStruct, Type: types.NewStructType(name)}
}
func NewIface(id int, name string) *Interface {
	return &Interface{ID: id, Name: name, Selector: GoInterface, Type: types.NewIfaceType(name)}
}
func NewCustom(id int, name string) *Custom {
	return &Custom{ID: id, Name: name, Selector: GoCustom, Type: types.NewCustomType(name)}
}
func NewMethod(id int, name string) *Method  { return &Method{ID: id, Name: name, Selector: GoMethod} }
func NewFunc(id int, name string) *Function  { return &Function{ID: id, Name: name, Selector: GoFunc} }
func NewConst(id int, name string) *Constant { return &Constant{ID: id, Name: name, Selector: GoConst} }
func NewVar(id int, name string) *Variable   { return &Variable{ID: id, Name: name, Selector: GoVar} }
func NewUnknown(id int, name string) *Unknown {
	return &Unknown{ID: id, Name: name, Selector: GoUnknown, Type: types.NewUnknownType(name)}
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
