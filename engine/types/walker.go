package types

import (
	"strconv"
	"strings"
)

func WalkType(def string) (Type, error) {
	if def == "" {
		return nil, nil
	}

	var hasPointer bool
	def = strings.TrimLeft(def, " ")
	if strings.HasPrefix(def, "*") {
		hasPointer = true
		def = strings.TrimLeft(def, "*")
	}
	var info = TypeInfo{def: def, pointer: hasPointer}

	switch IdentifyDescription(def) {
	case Map:
		info.descriptor = Map
		var m = &MapType{TypeInfo: info}

		key, residue := extract(def)
		k, err := WalkType(key)
		if err != nil {
			return nil, err
		}

		m.key = k
		m.subtype = Init(residue)

		return m, nil
	case Array:
		info.descriptor = Array
		var array = &ArrayType{TypeInfo: info}

		arrLen, residue := extract(def)
		i, _ := strconv.Atoi(arrLen)
		array.len = i
		array.subtype = Init(residue)

		return array, nil
	case Slice:
		info.descriptor = Slice
		var slice = &SliceType{TypeInfo: info}

		_, res := extract(def)
		slice.subtype = Init(res)

		return slice, nil
	case Chan:
		info.descriptor = Chan
		var ch = &ChanType{TypeInfo: info}

		val, res := extract(def)
		ch.ChanMode = chanMode(val)

		ch.subtype = Init(res)

		return ch, nil
	case Builtin:
		info.descriptor = Builtin
		return &BuiltinType{TypeInfo: info}, nil
	case Func:
		info.descriptor = Func
		return &FuncType{TypeInfo: info}, nil
	case EmptyInterface:
		info.descriptor = EmptyInterface
		return &EmptyInterfaceType{TypeInfo: info}, nil
	case EmptyStruct:
		info.descriptor = EmptyStruct
		return &EmptyStructType{TypeInfo: info}, nil
	case Imported:
		info.descriptor = Imported
		parts := strings.Split(info.def, ".")
		info.def = parts[1]
		return &ImportedType{TypeInfo: info, packName: parts[0]}, nil
	default:
		info.descriptor = Unknown
		return &UnknownType{TypeInfo: info}, nil
	}
}

var sum int

func lenTypes(t Type) int {
	if t == nil {
		return sum
	} else {
		sum += lenTypes(t.Subtype())
	}
	switch t.Descriptor() {
	case Struct, Interface, EmptyStruct, EmptyInterface, Func, Builtin:
		return sum + 1
	}
	return sum
}
