package types

const (
	Unknown TypeDescriptor = iota
	Package
	Import
	Imported
	Struct
	Interface
	EmptyInterface
	EmptyStruct
	Custom
	Builtin
	Func
	Chan
	Array
	Slice
	Map
)

func (d TypeDescriptor) String() string {
	var descriptors = map[TypeDescriptor]string{
		Unknown:        "UnknownType",
		Package:        "PackageType",
		Imported:       "ImportedType",
		Import:         "ImportType",
		Custom:         "CustomType",
		Struct:         "StructType",
		Interface:      "InterfaceType",
		EmptyInterface: "interface{}",
		EmptyStruct:    "struct{}",
		Builtin:        "BuiltinType",
		Func:           "FuncType",
		Chan:           "ChanType",
		Array:          "ArrayType",
		Slice:          "SliceType",
		Map:            "MapType",
	}
	return descriptors[d]
}
