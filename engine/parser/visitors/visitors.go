package visitors

import (
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/units"
)

type VisitMode int

const (
	Imports VisitMode = iota
	Interfaces
	Functions
	Customs
	Structs
	ValueDecls
)

var AllModes = []VisitMode{Imports, Interfaces, Functions, Customs, Structs, ValueDecls}

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
		Collection   collection.Collector
		modes        []VisitMode
		pack         units.Unit
		importsPaths []string
	}

	ImportVisitor struct {
		src          string
		Collector    collection.Collector
		pack         units.Unit
		importsPaths []string
	}

	StructsVisitor struct {
		parent     units.Unit //Transmitted only when another structure is embedded in the structure.
		structName string
		src        string
		pack       units.Unit
		Collector  collection.Collector
	}

	InterfacesVisitor struct {
		src       string
		ifaceName string
		pack      units.Unit
		parent    units.Unit //Transmitted only when another interface is embedded in the interface.
		Collector collection.Collector
		isField   bool
	}

	CustomsVisitor struct {
		src       string
		pack      units.Unit
		parent    units.Unit
		Collector collection.Collector
	}

	FuncsVisitor struct {
		src        string
		pack       units.Unit
		parent     units.Unit
		Collection collection.Collector
	}
)

func (vis *GeneralVisitor) FoundImports() []string { return vis.importsPaths }
