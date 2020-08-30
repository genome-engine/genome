package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
	"strings"
)

func (c *Collection) Add(root units.Unit, children ...units.Unit) error {
	c.log("Adding run. Number of adding elements: %d", len(children)+1)
	if root == nil {
		c.log("Error from collecting. Exit with error.")
		return fmt.Errorf("Passed root is empty - no addition possible. ")
	}
	//For logging informing.
	var adding string
	var addingCount int
	var add = func(unit units.Unit) {
		addingCount++
		adding += fmt.Sprintf("\t\t\t%d.{%v: %v}", addingCount, unit.GetSelector().Name(), unit.GetName())
		if addingCount > 1 {
			adding += "\n"
		}
	}

	//Getting root id.
	var key = root.GetId()

	if r, ok := c.rootTable[key]; !ok {
		c.rootTable[key] = root
		add(root)
	} else {
		if r.GetSelector() == units.GoUnknown && root.GetSelector() != units.GoUnknown {
			c.log("Replacing an unknown root with a senior root.")
			c.rootTable[key] = root
		}
	}

	for _, child := range children {
		if !root.GetSelector().CanContain(child.GetSelector()) {
			c.log("Relationship error detected.")
			return fmt.Errorf("%v can't contain %v. ", root.GetSelector().Name(), child.GetSelector().Name())
		}

		if children, ok := c.childrenTable[key]; ok && !UnitExist(children, child) {
			c.childrenTable[key] = append(c.childrenTable[key], child)
			add(child)
		} else if !ok {
			c.childrenTable[key] = append(c.childrenTable[key], child)
			add(child)
		}

		c.rootTable[child.GetId()] = child
	}

	c.log("Adding complete. Len elements:%d", len(c.rootTable))
	if adding != "" {
		c.log("Added units:\n%v", adding)
	} else {
		c.log("Added units: 0")
	}
	return nil
}

func (c *Collection) UnitsMap() map[units.Unit][]units.Unit {
	c.Clear()

	c.log("Prepare units map.")
	for id, unit := range c.rootTable {
		if child, ok := c.childrenTable[id]; ok {
			c.unitsMap[unit] = append(c.unitsMap[unit], child...)
			continue
		}

		c.unitsMap[unit] = nil
	}
	c.log("Prepare complete")
	return c.unitsMap
}

func (c *Collection) Merge(collector Collection) error {
	c.log("Merging collections")

	for root, children := range collector.UnitsMap() {
		err := c.Add(root, children...)
		if err != nil {
			c.log("Error from merging.Exit with error.")
			return err
		}
	}
	c.log("Merging complete.")
	return nil
}

func (c *Collection) Print(selectors ...units.Selector) {
	if len(selectors) == 0 {
		selectors = units.AllSelectors
	}

	for root, children := range c.UnitsMap() {
		if SelectorExist(selectors, root.GetSelector()) {
			fmt.Printf("%v{GetId: %v, out: %v}\n",
				root.GetSelector().Name(),
				root.GetId(),
				root.GetName(),
			)
			if len(children) == 0 {
				fmt.Printf("\t- No children\n")
			}
			for _, child := range children {
				fmt.Printf("\t- %v{GetId: %v, out: %v}\n",
					child.GetSelector().Name(),
					child.GetId(),
					child.GetName(),
				)
			}
			println()
		}
	}
}

func (c *Collection) findImplements() {
	c.log("Searching implements.")
	var implementsCount int
	var ifaceMethods = map[int][]units.Unit{}
	var ownersMethods = map[int][]units.Unit{}

	for _, unit := range c.rootTable {
		switch unit.GetSelector() {
		case units.GoInterface:
			for _, method := range c.childrenTable[unit.GetId()] {
				if method.GetSelector() == units.GoMethod {
					ifaceMethods[unit.GetId()] = append(ifaceMethods[unit.GetId()], method)
				}
			}
		case units.GoStruct, units.GoCustom:
			for _, method := range c.childrenTable[unit.GetId()] {
				if method.GetSelector() == units.GoMethod {
					ownersMethods[unit.GetId()] = append(ownersMethods[unit.GetId()], method)
				}
			}
		}
	}

	for ifaceId, iMethods := range ifaceMethods {
		for ownerId, oMethods := range ownersMethods {
			if compareMethods(iMethods, oMethods) {
				implementsCount++
				owner := c.rootTable[ownerId]
				iface := c.rootTable[ifaceId]
				c.childrenTable[owner.GetId()] = append(c.childrenTable[owner.GetId()], iface)
			}
		}
	}
	c.log("Search complete.Number of implementations found: %d", implementsCount)
}

func (c *Collection) Linking() {
	c.log("Linking start.")
	for _, unit := range c.rootTable {
		switch unit.GetSelector() {
		case units.GoConst:
			con := unit.(*units.Const)
			if con.Enum && con.Type != "int" {
				for _, custom := range c.rootTable {
					if custom.GetSelector() == units.GoCustom && strings.Contains(con.Type, custom.GetName()) {
						var customMethods []units.Unit
						c.log("Enum custom type was found.")
						c.childrenTable[custom.GetId()] = append(c.childrenTable[custom.GetId()], con)
						for _, method := range c.childrenTable[custom.GetId()] {
							if m, ok := method.(*units.Method); ok {
								customMethods = append(customMethods, m)
							}
						}
						c.childrenTable[con.GetId()] = append(c.childrenTable[con.GetId()], customMethods...)
					}
				}
			}
		}
	}
	c.findImplements()
	c.log("Linking end.")
}

func (c *Collection) Clear() {
	c.log("Clearing collection.")
	c.unitsMap = map[units.Unit][]units.Unit{}
	c.log("Clearing end.")
}

func compareMethods(ifaceMethods, ownerMethods []units.Unit) bool {
	if len(ifaceMethods) > len(ownerMethods) {
		return false
	}
	var (
		necessaryMatches = len(ifaceMethods)
		actualMatches    int
	)

	for _, ifaceMethod := range ifaceMethods {
		if UnitExist(ownerMethods, ifaceMethod) {
			actualMatches++
		}
	}

	return necessaryMatches == actualMatches
}
