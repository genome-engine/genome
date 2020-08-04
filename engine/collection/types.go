package collection

import (
	"github.com/genome-engine/genome/engine/units"
)

type (
	//UnitsSearcher is used in the templating engine.
	UnitsSearcher interface {
		Search(selectors ...units.Selector) []units.Unit
		SearchChildren(unit units.Unit, selectors ...units.Selector) ([]units.Unit, error)
		SearchParents(unit units.Unit, selector ...units.Selector) ([]units.Unit, error)
		SearchById(id int) units.Unit
	}

	//UnitsManager is used in parser.
	UnitsManager interface {
		Add(root units.Unit, children ...units.Unit) error
		Merge(collector Collector) error
		Linking()
		Clear()
	}

	//The Collector simply combines the UnitsSearcher and the UnitsManager
	Collector interface {
		UnitsMap() map[units.Unit][]units.Unit
		Print(selectors ...units.Selector)
		UnitsManager
		UnitsSearcher
	}

	Collection struct {
		unitsMap      UnitsMap
		rootTable     map[int]units.Unit
		childrenTable map[int][]units.Unit
	}

	UnitsMap map[units.Unit][]units.Unit
)

func New() *Collection {
	return &Collection{
		unitsMap:      map[units.Unit][]units.Unit{},
		rootTable:     map[int]units.Unit{},
		childrenTable: map[int][]units.Unit{},
	}
}
