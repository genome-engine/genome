package units

type GoSelector int

//Default selectors relevant in the author's opinion
const (
	GoUnknown GoSelector = iota
	GoInterface
	GoPackage
	GoImport
	GoStruct
	GoCustom
	GoMethod
	GoConst
	GoFunc
	GoVar
)

//The names are used for further output to the console.
var selectors = map[GoSelector]string{
	GoInterface: "Interface",
	GoUnknown:   "Unknown",
	GoPackage:   "Package",
	GoStruct:    "Struct",
	GoCustom:    "Custom", //Subtype out {{no struct}}
	GoImport:    "Import",
	GoConst:     "Const",
	GoMethod:    "Method", //func(owner) out(){}
	GoFunc:      "Func",
	GoVar:       "Var",
}

var AllSelectors = []Selector{
	GoUnknown, GoInterface, GoFunc, GoVar, GoConst, GoPackage, GoMethod, GoCustom, GoStruct, GoImport,
}

var (
	PackChildren = []Selector{
		GoInterface, GoMethod, GoImport, GoStruct, GoCustom, GoFunc, GoConst, GoVar,
	}
	MethodChildren  = []Selector{GoConst, GoVar, GoFunc, GoCustom}
	FuncChildren    = []Selector{GoConst, GoVar, GoFunc, GoCustom}
	StructChildren  = []Selector{GoStruct, GoInterface, GoMethod, GoCustom}
	CustomChildren  = []Selector{GoMethod, GoInterface, GoConst}
	IfaceChildren   = []Selector{GoMethod, GoInterface}
	UnknownChildren = []Selector{GoMethod}
)

//The default allowable nesting boxes are shown here.
var possibleContains = map[GoSelector][]Selector{
	GoPackage: PackChildren, GoMethod: MethodChildren, GoFunc: FuncChildren, GoStruct: StructChildren,
	GoCustom: CustomChildren, GoInterface: IfaceChildren,
	GoUnknown: UnknownChildren,
	GoConst:   {GoMethod},
	//Can't contain anything.
	GoImport: nil, GoVar: nil,
}

//The Package method is equivalent to the String method.
//But if the Selector is not known, it will return Unknown.
func (s GoSelector) Name() string {
	if typ, ok := selectors[s]; ok {
		return typ
	}
	return selectors[GoUnknown]
}

//Checks whether the object with the transmitted Selector can be nested in the default_template object.
func (s GoSelector) CanContain(selector Selector) bool {
	possible, ok := possibleContains[s]
	if !ok || possible == nil {
		return false
	}

	for _, goSelector := range possible {
		if goSelector == selector {
			return true
		}
	}

	return false
}

func (s GoSelector) ChildSelectors() []Selector {
	switch s {
	case GoPackage:
		return PackChildren
	case GoMethod:
		return MethodChildren
	case GoStruct:
		return StructChildren
	case GoCustom:
		return CustomChildren
	case GoFunc:
		return FuncChildren
	case GoInterface:
		return IfaceChildren
	case GoUnknown:
		return UnknownChildren
	default:
		return nil
	}
}

func (s GoSelector) ParentSelectors() []Selector {
	var parents []Selector

	for selector, children := range possibleContains {
		for _, child := range children {
			if child == s {
				parents = append(parents, selector)
			}
		}
	}

	return parents
}

func ToSelector(str string) Selector {
	for selector, name := range selectors {
		if str == name {
			return selector
		}
	}
	return GoUnknown
}
