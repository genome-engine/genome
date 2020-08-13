package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
	"strings"
)

func (c *Collection) Clear() {
	c.unitsMap = map[units.Unit][]units.Unit{}
}

func (c *Collection) Add(root units.Unit, children ...units.Unit) error {
	if root == nil {
		return fmt.Errorf("Passed root is empty - no addition possible. ")
	}

	var key = root.GetId()

	if r, ok := c.rootTable[key]; !ok {
		c.rootTable[key] = root
	} else {
		if r.GetSelector() == units.GoUnknown && root.GetSelector() != units.GoUnknown {
			c.rootTable[key] = root
		}
	}

	for _, child := range children {
		if !root.GetSelector().CanContain(child.GetSelector()) {
			return fmt.Errorf("%v can't contain %v. ", root.GetSelector().Name(), child.GetSelector().Name())
		}

		if children, ok := c.childrenTable[key]; ok && !UnitExist(children, child) {
			c.childrenTable[key] = append(c.childrenTable[key], child)
		} else if !ok {
			c.childrenTable[key] = append(c.childrenTable[key], child)
		}

		c.rootTable[child.GetId()] = child
	}

	return nil
}

func (c *Collection) UnitsMap() map[units.Unit][]units.Unit {
	c.Clear()

	for id, unit := range c.rootTable {
		if child, ok := c.childrenTable[id]; ok {
			c.unitsMap[unit] = append(c.unitsMap[unit], child...)
			continue
		}

		c.unitsMap[unit] = nil
	}

	return c.unitsMap
}

func (c *Collection) Merge(collector Collector) error {
	if collector == nil {
		return nil
	}

	for root, children := range collector.UnitsMap() {
		err := c.Add(root, children...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection) Print(selectors ...units.Selector) {
	if len(selectors) == 0 {
		selectors = units.AllSelectors
	}

	for root, children := range c.UnitsMap() {
		if SelectorExist(selectors, root.GetSelector()) {
			fmt.Printf("%v{GetId: %v, Name: %v}\n",
				root.GetSelector().Name(),
				root.GetId(),
				root.GetName(),
			)
			if len(children) == 0 {
				fmt.Printf("\t- No children\n")
			}
			for _, child := range children {
				fmt.Printf("\t- %v{GetId: %v, Name: %v}\n",
					child.GetSelector().Name(),
					child.GetId(),
					child.GetName(),
				)
			}
			println()
		}
	}
}

func (c *Collection) Linking() {
	for _, unit := range c.rootTable {
		switch unit.GetSelector() {
		case units.GoConst:
			con := unit.(*units.Constant)
			if con.Enum && con.Type != "int" {
				for _, custom := range c.rootTable {
					if custom.GetSelector() == units.GoCustom && strings.Contains(con.Type, custom.GetName()) {
						var customMethods []units.Unit
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
}

func (c *Collection) findImplements() {
	var ifaceMethods = map[int][]units.Unit{}
	var ownersMetods = map[int][]units.Unit{}

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
					ownersMetods[unit.GetId()] = append(ownersMetods[unit.GetId()], method)
				}
			}
		}
	}

	for ifaceId, iMethods := range ifaceMethods {
		for ownerId, oMethods := range ownersMetods {
			if compareMethods(iMethods, oMethods) {
				owner := c.rootTable[ownerId]
				iface := c.rootTable[ifaceId]
				c.childrenTable[owner.GetId()] = append(c.childrenTable[owner.GetId()], iface)
			}
		}
	}
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
