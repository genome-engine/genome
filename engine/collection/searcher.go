package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
)

//This function allows you to find units with specific selectors. It is searched at both root and child levels.
func (c *Collection) Search(selectors ...units.Selector) (u []units.Unit) {
	c.log("Search by %d selectors", len(selectors))
	if len(selectors) == 0 {
		selectors = units.AllSelectors
	}
	var selectorsMap = map[string]bool{}
	//It's important. UnitsMap is accessed only through the UnitsMap function,
	//because it is at the time of its call that root->children are linked.
	for _, root := range c.rootTable {
		if SelectorExist(selectors, root.GetSelector()) && !UnitExist(u, root) {
			selectorsMap[root.GetSelector().Name()] = true
			u = append(u, root)
			continue
		}
	}

	c.log("Search complete. Found units: %d.", len(u))
	for s, b := range selectorsMap {
		if b {
			c.log("Units with %v selector was found.", s)
		}
	}
	return u
}

//This function allows you to find units with specific id. It is searched at both root and child levels.
func (c *Collection) SearchById(id int) units.Unit {
	c.log("Search unit by id:%d.", id)
	//It's important. UnitsMap is accessed only through the UnitsMap function,
	//because it is at the time of its call that root->children are linked.
	for root, children := range c.UnitsMap() {
		//root level
		if root.GetId() == id {
			return root
		}

		//child level
		for _, child := range children {
			if child.GetId() == id {
				c.log("Unit was found.")
				return child
			}
		}
	}

	c.log("Unit not found.")
	return nil
}

//The child search method searches for nested objects.
//Returns an error if wrong selectors are passed.
func (c *Collection) SearchChildren(unit units.Unit, selectors ...units.Selector) (children []units.Unit, err error) {
	c.log("Searching children by root.id:%d.", unit.GetId())
	objects, ok := c.childrenTable[unit.GetId()]

	var selectorsMap = map[string]bool{}
	//Is the object on the map?
	if !ok {
		c.log("Not records with this id. Exit with error.")
		return nil, fmt.Errorf("An object with an id %v is not registered in the address book. ", unit.GetId())
	}

	//If there are no selectors, it will return all the child objects
	if len(selectors) == 0 {
		return objects, nil
	}

	//Testing for selector compatibility.
	for _, selector := range selectors {
		if !unit.GetSelector().CanContain(selector) {
			c.log("Relationship error detected.")
			return nil, fmt.Errorf(
				"%v cannot contain %v, so no search is possible. ",
				unit.GetSelector().Name(), selector.Name())
		}
	}

	for _, info := range objects {
		//If the child object's selector matches the ones that were passed, the object will be added.
		if SelectorExist(selectors, info.GetSelector()) {
			c.log("Child found.")
			selectorsMap[info.GetSelector().Name()] = true
			children = append(children, info)
		}
	}

	c.log("Searching complete. Number of children: %d", len(children))
	for s, b := range selectorsMap {
		if b {
			c.log("Child with %v selector was found.", s)
		}
	}
	return children, nil
}

//Parental object search searches for top level objects that contain the object you are searching for.
//Returns an error if wrong selectors are passed.
func (c *Collection) SearchParents(unit units.Unit, selectors ...units.Selector) (roots []units.Unit, err error) {
	var selectorsMap = map[string]bool{}
	c.log("Parents searching by root.id:%d", unit.GetId())
	//Checking if the selector of the transmitted unit can be nested in selectors from the selectors list.
	for _, selector := range selectors {
		c.log("Relationship error was detected.")
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
			selectorsMap[root.GetSelector().Name()] = true
			roots = append(roots, root)
		}
	}

	c.log("Searching complete.Number of parents: %d", len(roots))
	for s, b := range selectorsMap {
		if b {
			c.log("Child with %v selector was found.", s)
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
