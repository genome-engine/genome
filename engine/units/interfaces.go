package units

import "github.com/genome-engine/genome/engine/types"

type (
	//The Unit interface is used to universalize elements of the collection.
	//This data will then have to go through type conversion to basic structures.
	Unit interface {
		GetId() int
		GetName() string
		GetType() types.Type
		GetSelector() Selector
		SetType(p types.Type)
	}

	//The Selector is created for identification of an accessory of default_template object
	//and also for creation of possible variants of nesting of selectors in each other.
	//More details about selectors and units are written in ./units/README.md,
	//i.e. in the basic implementation of units and selectors.
	Selector interface {
		Name() string
		CanContain(selector Selector) bool
		ChildSelectors() []Selector
		ParentSelectors() []Selector
	}
)
