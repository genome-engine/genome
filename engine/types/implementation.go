package types

func (t *StructType) Definition() string         { return t.def }
func (t *StructType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *StructType) Subtype() Type              { return t.subtype }
func (t *StructType) LenTypes() int              { return 0 }
func (t *StructType) SetSubtype(Type)            { return } //the structure may not contain subtypes

func (t *EmptyStructType) Definition() string         { return t.def }
func (t *EmptyStructType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *EmptyStructType) Subtype() Type              { return t.subtype }
func (t *EmptyStructType) LenTypes() int              { return 0 }
func (t *EmptyStructType) SetSubtype(Type)            { return } //the empty structure may not contain subtypes

func (t *InterfaceType) Definition() string         { return t.def }
func (t *InterfaceType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *InterfaceType) Subtype() Type              { return t.subtype }
func (t *InterfaceType) LenTypes() int              { return 0 }
func (t *InterfaceType) SetSubtype(Type)            { return } //the interface may not contain subtypes

func (t *EmptyInterfaceType) Definition() string         { return t.def }
func (t *EmptyInterfaceType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *EmptyInterfaceType) Subtype() Type              { return t.subtype }
func (t *EmptyInterfaceType) LenTypes() int              { return 0 }
func (t *EmptyInterfaceType) SetSubtype(Type)            { return } //the empty interface may not contain subtypes

func (t *BuiltinType) Definition() string         { return t.def }
func (t *BuiltinType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *BuiltinType) Subtype() Type              { return t.subtype }
func (t *BuiltinType) LenTypes() int              { return 0 }
func (t *BuiltinType) SetSubtype(Type)            { return } //the builtin may not contain subtypes

func (t *FuncType) Definition() string         { return t.def }
func (t *FuncType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *FuncType) Subtype() Type              { return t.subtype }
func (t *FuncType) LenTypes() int              { return 0 }
func (t *FuncType) SetSubtype(Type)            { return } //the func may not contain subtypes

func (t *ChanType) Definition() string         { return t.def }
func (t *ChanType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *ChanType) Subtype() Type              { return t.subtype }
func (t *ChanType) LenTypes() int              { return lenTypes(t) }
func (t *ChanType) SetSubtype(sub Type)        { t.subtype = sub }

func (t *ArrayType) Definition() string         { return t.def }
func (t *ArrayType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *ArrayType) Subtype() Type              { return t.subtype }
func (t *ArrayType) LenTypes() int              { return lenTypes(t) }
func (t *ArrayType) SetSubtype(sub Type)        { t.subtype = sub }
func (t *ArrayType) Len() int                   { return t.len } //non-interface

func (t *SliceType) Definition() string         { return t.def }
func (t *SliceType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *SliceType) Subtype() Type              { return t.subtype }
func (t *SliceType) LenTypes() int              { return lenTypes(t) }
func (t *SliceType) SetSubtype(sub Type)        { t.subtype = sub }

func (t *MapType) Definition() string         { return t.def }
func (t *MapType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *MapType) Subtype() Type              { return t.subtype }
func (t *MapType) LenTypes() int              { return lenTypes(t) }
func (t *MapType) SetSubtype(sub Type)        { t.subtype = sub }
func (t *MapType) Key() Type                  { return t.key } //non-interface

func (t *UnknownType) Definition() string         { return t.def }
func (t *UnknownType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *UnknownType) Subtype() Type              { return t.subtype }
func (t *UnknownType) LenTypes() int              { return lenTypes(t) }
func (t *UnknownType) SetSubtype(sub Type)        { t.subtype = sub }

func (t *PackageType) Definition() string         { return t.def }
func (t *PackageType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *PackageType) Subtype() Type              { return t.subtype }
func (t *PackageType) LenTypes() int              { return lenTypes(t) }
func (t *PackageType) SetSubtype(Type)            { return } //the package may not contain subtypes

func (t *ImportType) Definition() string         { return t.def }
func (t *ImportType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *ImportType) Subtype() Type              { return t.subtype }
func (t *ImportType) LenTypes() int              { return lenTypes(t) }
func (t *ImportType) SetSubtype(Type)            { return } //the import may not contain subtypes

func (t *ImportedType) Definition() string         { return t.def }
func (t *ImportedType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *ImportedType) Subtype() Type              { return t.subtype }
func (t *ImportedType) LenTypes() int              { return lenTypes(t) }
func (t *ImportedType) SetSubtype(Type)            { return }            //the imported may not contain subtypes
func (t *ImportedType) PackName() string           { return t.packName } //non-interface

func (t *CustomType) Definition() string         { return t.def }
func (t *CustomType) Descriptor() TypeDescriptor { return t.descriptor }
func (t *CustomType) Subtype() Type              { return t.subtype }
func (t *CustomType) LenTypes() int              { return lenTypes(t) }
func (t *CustomType) SetSubtype(Type)            { return } //the imported may not contain subtypes
