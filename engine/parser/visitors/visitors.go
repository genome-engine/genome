package visitors

import (
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/units"
)

type VisitMode int

const (
	All VisitMode = iota
	Imports
	Interfaces
	Functions
	Customs
	Structs
	Values
)

var modes = map[VisitMode]string{
	All:        "all",
	Imports:    "imports",
	Interfaces: "ifaces",
	Functions:  "funcs",
	Customs:    "customs",
	Structs:    "structs",
	Values:     "values",
}

func (m VisitMode) String() string {
	mode, ok := modes[m]
	if ok {
		return mode
	}
	return ""
}

func ToMode(s string) VisitMode {
	for mode, name := range modes {
		if name == s {
			return mode
		}
	}
	return Values
}

var AllModes = []VisitMode{Imports, Interfaces, Functions, Customs, Structs, Values}

func modeExist(modes []VisitMode, mode VisitMode) bool {
	for _, m := range modes {
		if m == mode {
			return true
		}
	}
	return false
}

type (
	GeneralVisitor struct {
		path         string
		packMainDir  string
		src          string
		Collection   collection.Collection
		modes        []VisitMode
		pack         units.Unit
		importsPaths []string
	}

	ImportVisitor struct {
		src          string
		Collector    collection.Collection
		pack         units.Unit
		importsPaths []string
	}

	StructsVisitor struct {
		parent     units.Unit //Transmitted only when another structure is embedded in the structure.
		structName string
		src        string
		comment    string
		pack       units.Unit
		Collector  collection.Collection
	}

	InterfacesVisitor struct {
		src       string
		ifaceName string
		comment   string
		pack      units.Unit
		parent    units.Unit //Transmitted only when another interface is embedded in the interface.
		Collector collection.Collection
		isField   bool
	}

	CustomsVisitor struct {
		src       string
		comment   string
		pack      units.Unit
		parent    units.Unit
		Collector collection.Collection
	}

	FuncsVisitor struct {
		src        string
		pack       units.Unit
		parent     units.Unit
		Collection collection.Collection
	}
)

func (vis *GeneralVisitor) FoundImports() []string { return vis.importsPaths }
