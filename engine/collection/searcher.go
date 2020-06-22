package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
)

//This function allows you to find units with specific selectors. It is searched at both root and child levels.
func (c *Collection) SearchBySelectors(selectors ...units.Selector) (u []units.Unit) {
	//It's important. ObjectMap is accessed only through the GetObjectMap function,
	//because it is at the time of its call that root->children are linked.
	for root, childes := range c.GetObjectMap() {
		//root level
		if SelectorExist(selectors, root.GetSelector()) && !UnitExist(u, root) {
			u = append(u, root)
			continue
		}

		for _, child := range childes {
			if SelectorExist(selectors, child.GetSelector()) && !UnitExist(u, child) {
				u = append(u, child)
			}
		}
	}

	return u
}

//This function allows you to find units with specific id. It is searched at both root and child levels.
func (c *Collection) SearchById(id int) units.Unit {
	//It's important. ObjectMap is accessed only through the GetObjectMap function,
	//because it is at the time of its call that root->children are linked.
	for root, children := range c.GetObjectMap() {
		//root level
		if root.GetId() == id {
			return root
		}

		//child level
		for _, child := range children {
			if child.GetId() == id {
				return child
			}
		}
	}
	return nil
}

//The child search method searches for nested objects.
//Returns an error if wrong selectors are passed.
func (c *Collection) SearchChildren(unit units.Unit, selectors ...units.Selector) (children []units.Unit, err error) {
	objects, ok := c.childrenTable[unit.GetId()]

	//Is the object on the map?
	if !ok {
		return nil, fmt.Errorf("An object with an id %v is not registered in the address book. ", unit.GetId())
	}

	//If there are no selectors, it will return all the child objects
	if len(selectors) == 0 {
		return objects, nil
	}

	//Testing for selector compatibility.
	for _, selector := range selectors {
		if !unit.GetSelector().CanContain(selector) {
			return nil, fmt.Errorf(
				"%v cannot contain %v, so no search is possible. ",
				unit.GetSelector().Name(), selector.Name())
		}
	}

	for _, info := range objects {
		//If the child object's selector matches the ones that were passed, the object will be added.
		if SelectorExist(selectors, info.GetSelector()) {
			children = append(children, info)
		}
	}

	return children, nil
}

//Parental object search searches for top level objects that contain the object you are searching for.
//Returns an error if wrong selectors are passed.
func (c *Collection) SearchParents(unit units.Unit, selectors ...units.Selector) (roots []units.Unit, err error) {
	//Checking if the selector of the transmitted unit can be nested in selectors from the selectors list.
	for _, selector := range selectors {
		if !selector.CanContain(unit.GetSelector()) {
			return nil, fmt.Errorf(
				"%v cannot contain %v, so no search is possible. ",
				unit.GetSelector().Name(), selector.Name())
		}
	}

	//Passing through all the units in rootTable
	for _, root := range c.rootTable {
		//If a unit is contained in childrenTable with the root key,
		//this root is added to the roots, which will then be returned to
		if UnitExist(c.childrenTable[root.GetId()], unit) {
			roots = append(roots, root)
		}
	}

	return roots, nil
}

//Auxiliary function for finding a selector in a selector array.
func SelectorExist(selectors []units.Selector, selector units.Selector) bool {
	for _, d := range selectors {
		if d == selector {
			return true
		}
	}
	return false
}

func UnitExist(units []units.Unit, unit units.Unit) bool {
	for _, u := range units {
		if u.GetId() == unit.GetId() {
			return true
		}
	}
	return false
}
