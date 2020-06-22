package types

import (
	"fmt"
	"strings"
)

const (
	EmptyStructure = "struct{}"
	EmptyIface     = "interface{}"
)

func NewPackageType(def string) *PackageType {
	return &PackageType{TypeInfo: TypeInfo{def: def, descriptor: Package}}
}
func NewImportType(def string) *ImportType {
	return &ImportType{TypeInfo: TypeInfo{def: def, descriptor: Import}}
}
func NewImportedType(def string) (*ImportedType, error) {
	if !isImported(def) {
		return nil, fmt.Errorf("%v type does not contain an imported element at the beginning", def)
	}

	typParts := strings.Split(def, ".")
	return &ImportedType{TypeInfo: TypeInfo{def: typParts[1], descriptor: Imported}, packName: typParts[0]}, nil
}
func NewStructType(def string) *StructType {
	return &StructType{TypeInfo: TypeInfo{def: def, descriptor: Struct}}
}
func NewEmptyStructType() *EmptyStructType {
	return &EmptyStructType{TypeInfo: TypeInfo{def: "struct{}", descriptor: EmptyStruct}}
}
func NewIfaceType(def string) *InterfaceType {
	return &InterfaceType{TypeInfo: TypeInfo{def: def, descriptor: Interface}}
}
func NewEmptyInterfaceType() *EmptyInterfaceType {
	return &EmptyInterfaceType{TypeInfo: TypeInfo{def: "interface{}", descriptor: EmptyInterface}}
}
func NewCustomType(def string) *CustomType {
	return &CustomType{TypeInfo: TypeInfo{def: def, descriptor: Custom}}
}
func NewBuiltinType(def string) (*BuiltinType, error) {
	if !isBuiltin(def) {
		return nil, fmt.Errorf("%v type does not contain an builtion element at the beginning", def)
	}
	info := TypeInfo{def: def, descriptor: Builtin}
	return &BuiltinType{TypeInfo: info}, nil
}
func NewFuncType(def string) (*FuncType, error) {
	if !isFunc(def) {
		return nil, fmt.Errorf("%v type does not contain an func element at the beginning", def)
	}
	def = strings.TrimLeft(def, "func ")
	info := TypeInfo{def: def, descriptor: Func}
	return &FuncType{TypeInfo: info}, nil
}
func NewChanType(def string) (*ChanType, error) {
	if !isChan(def) {
		return nil, fmt.Errorf("%v type does not contain an chan element at the beginning", def)
	}

	_, t := extract(def)

	typ, err := WalkType(t)

	if err != nil {
		return nil, err
	}

	info := TypeInfo{def: def, descriptor: Chan, subtype: typ}
	return &ChanType{TypeInfo: info}, nil
}
func NewArrayType(def string) (*ArrayType, error) {
	if !isArray(def) {
		return nil, fmt.Errorf("%v type does not contain an array element at the beginning", def)
	}

	_, t := extract(def)

	typ, err := WalkType(t)
	if err != nil {
		return nil, err
	}
	info := TypeInfo{def: def, descriptor: Array, subtype: typ}
	return &ArrayType{TypeInfo: info}, nil
}
func NewSliceType(def string) (*SliceType, error) {
	if !isSlice(def) {
		return nil, fmt.Errorf("%v type does not contain an slice element at the beginning", def)
	}
	_, t := extract(def)

	typ, err := WalkType(t)
	if err != nil {
		return nil, err
	}

	info := TypeInfo{def: def, descriptor: Slice, subtype: typ}
	return &SliceType{TypeInfo: info}, nil
}
func NewMapType(def string) (*MapType, error) {
	if !isMap(def) {
		return nil, fmt.Errorf("%v type does not contain an array element at the beginning", def)
	}

	k, t := extract(def)

	typ, err := WalkType(t)
	if err != nil {
		return nil, err
	}

	key, err := WalkType(k)
	if err != nil {
		return nil, err
	}

	info := TypeInfo{def: def, descriptor: Map, subtype: typ}
	return &MapType{TypeInfo: info, key: key}, nil
}
func NewUnknownType(definition string) *UnknownType {
	return &UnknownType{TypeInfo: TypeInfo{def: definition, descriptor: Unknown}}
}

func Init(def string) Type {
	description := IdentifyDescription(def)

	switch description {
	case EmptyStruct:
		return NewEmptyStructType()
	case EmptyInterface:
		return NewEmptyInterfaceType()
	case Imported:
		i, _ := NewImportedType(def)
		return i
	case Builtin:
		b, _ := NewBuiltinType(def)
		return b
	case Func:
		f, _ := NewFuncType(def)
		return f
	case Chan:
		ch, _ := NewChanType(def)
		return ch
	case Array:
		a, _ := NewArrayType(def)
		return a
	case Slice:
		s, _ := NewSliceType(def)
		return s
	case Map:
		m, _ := NewMapType(def)
		return m
	default:
		return NewUnknownType(def)
	}
}
