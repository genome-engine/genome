package units

func NewPack(id int, name string) *Pack {
	return &Pack{ID: id, Name: name, Selector: GoPackage}
}
func NewImport(id int, name string) *Import { return &Import{ID: id, Name: name, Selector: GoImport} }
func NewStruct(id int, name string) *Struct {
	return &Struct{ID: id, Name: name, Selector: GoStruct}
}
func NewIface(id int, name string) *Iface {
	return &Iface{ID: id, Name: name, Selector: GoInterface}
}
func NewCustom(id int, name string) *Custom { return &Custom{ID: id, Name: name, Selector: GoCustom} }
func NewMethod(id int, name string) *Method { return &Method{ID: id, Name: name, Selector: GoMethod} }
func NewFunc(id int, name string) *Func     { return &Func{ID: id, Name: name, Selector: GoFunc} }
func NewConst(id int, name string) *Const   { return &Const{ID: id, Name: name, Selector: GoConst} }
func NewVar(id int, name string) *Var       { return &Var{ID: id, Name: name, Selector: GoVar} }
func NewUnknown(id int, name string) *Unknown {
	return &Unknown{ID: id, Name: name, Selector: GoUnknown}
}

func Init(id int, name string, selector Selector) Unit {
	switch selector {
	case GoPackage:
		return NewPack(id, name)
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
