package types

type (
	TypeDescriptor int

	Type interface {
		Definition() string
		Descriptor() TypeDescriptor
		Subtype() Type
		LenTypes() int
		SetSubtype(subtype Type)
	}

	TypeInfo struct {
		def        string
		descriptor TypeDescriptor
		subtype    Type
		pointer    bool
	}

	PackageType        struct{ TypeInfo }
	ImportType         struct{ TypeInfo }
	StructType         struct{ TypeInfo }
	EmptyStructType    struct{ TypeInfo }
	InterfaceType      struct{ TypeInfo }
	EmptyInterfaceType struct{ TypeInfo }
	CustomType         struct{ TypeInfo }
	BuiltinType        struct{ TypeInfo }
	FuncType           struct{ TypeInfo }
	SliceType          struct{ TypeInfo }
	UnknownType        struct{ TypeInfo }
	ChanType           struct {
		TypeInfo
		ChanMode
	}
	ArrayType struct {
		TypeInfo
		len int
	}
	MapType struct {
		TypeInfo
		key Type
	}
	ImportedType struct {
		TypeInfo
		packName string
	}
)
