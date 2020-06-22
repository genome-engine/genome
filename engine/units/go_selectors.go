package units

type GoSelector int

//Default selectors relevant in the author's opinion
const (
	GoUnknown GoSelector = iota
	GoPackage
	GoImport
	GoStruct
	GoCustom
	GoInterface
	GoMethod
	GoFunc
	GoEnumSeries
	GoConst
	GoVar
)

//The names are used for further output to the console.
var selectors = map[GoSelector]string{
	GoInterface:  "Interface",
	GoUnknown:    "Unknown",
	GoPackage:    "Package",
	GoStruct:     "Struct",
	GoCustom:     "Custom", //Subtype name {{no struct}}
	GoImport:     "Import",
	GoEnumSeries: "EnumSeries",
	GoConst:      "Const",
	GoMethod:     "Method", //func(owner) name(){}
	GoFunc:       "Func",
	GoVar:        "Var",
}

var AllSelectors = []Selector{
	GoUnknown, GoInterface, GoFunc, GoVar, GoConst, GoPackage, GoMethod, GoCustom, GoStruct, GoImport, GoEnumSeries,
}

var (
	packPossibles = []GoSelector{
		GoInterface, GoMethod, GoImport, GoStruct, GoCustom, GoFunc, GoConst, GoVar, GoEnumSeries,
	}
	methodPossibles     = []GoSelector{GoConst, GoVar, GoFunc, GoCustom}
	funcPossibles       = []GoSelector{GoConst, GoVar, GoFunc, GoCustom}
	structPossibles     = []GoSelector{GoStruct, GoInterface, GoMethod, GoCustom}
	customPossibles     = []GoSelector{GoMethod, GoInterface}
	ifacePossibles      = []GoSelector{GoMethod, GoInterface}
	unknownPossibles    = []GoSelector{GoMethod}
	enumSeriesPossibles = []GoSelector{GoConst}
)

//The default allowable nesting boxes are shown here.
var possibleContains = map[GoSelector][]GoSelector{
	GoPackage: packPossibles, GoMethod: methodPossibles, GoFunc: funcPossibles, GoStruct: structPossibles,
	GoCustom: customPossibles, GoInterface: ifacePossibles,
	GoUnknown: unknownPossibles, GoEnumSeries: enumSeriesPossibles,
	//Can't contain anything.
	GoConst: nil, GoImport: nil, GoVar: nil,
}

//The Package method is equivalent to the String method.
//But if the selector is not known, it will return Unknown.
func (s GoSelector) Name() string {
	if typ, ok := selectors[s]; ok {
		return typ
	}
	return selectors[GoUnknown]
}

//Checks whether the object with the transmitted selector can be nested in the default_template object.
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

func ToSelector(str string) Selector {
	for selector, name := range selectors {
		if str == name {
			return selector
		}
	}
	return GoUnknown
}
