package collection

import (
	"github.com/genome-engine/genome/engine/units"
)

type (
	//UnitsSearcher is used in the templating engine.
	UnitsSearcher interface {
		SearchBySelectors(selectors ...units.Selector) []units.Unit
		SearchChildren(unit units.Unit, selectors ...units.Selector) ([]units.Unit, error)
		SearchParents(unit units.Unit, selector ...units.Selector) ([]units.Unit, error)
		SearchById(id int) units.Unit
	}

	//UnitsManager is used in parser.
	UnitsManager interface {
		Add(root units.Unit, children ...units.Unit) error
		JoinCollection(collector Collector) error
		Clear()
	}

	//The Collector simply combines the UnitsSearcher and the UnitsManager
	Collector interface {
		GetObjectMap() map[units.Unit][]units.Unit
		Print(selectors ...units.Selector)
		UnitsManager
		UnitsSearcher
	}

	Collection struct {
		mode          AddonMode
		objectMap     ObjectMap
		rootTable     map[int]units.Unit
		childrenTable map[int][]units.Unit
	}

	AddonMode int
	ObjectMap map[units.Unit][]units.Unit
)

func New(mode AddonMode) *Collection {
	return &Collection{
		objectMap:     map[units.Unit][]units.Unit{},
		rootTable:     map[int]units.Unit{},
		childrenTable: map[int][]units.Unit{},
		mode:          mode,
	}
}

const (
	WithChildless AddonMode = iota
	WithoutChildless
)
