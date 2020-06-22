package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
)

func (c *Collection) Clear() {
	c.objectMap = map[units.Unit][]units.Unit{}
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

		if _, ok := c.rootTable[child.GetId()]; !ok && c.mode == WithChildless {
			c.rootTable[child.GetId()] = child
		}
	}

	return nil
}

func (c *Collection) GetObjectMap() map[units.Unit][]units.Unit {
	c.Clear()

	for id, unit := range c.rootTable {
		if child, ok := c.childrenTable[id]; ok {
			c.objectMap[unit] = append(c.objectMap[unit], child...)
			continue
		}

		if c.mode == WithChildless {
			c.objectMap[unit] = nil
		}
	}

	return c.objectMap
}

func (c *Collection) JoinCollection(collector Collector) error {
	if collector == nil {
		return fmt.Errorf("You handed over an empty collection ")
	}

	for root, children := range collector.GetObjectMap() {
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

	for root, children := range c.GetObjectMap() {
		if SelectorExist(selectors, root.GetSelector()) {
			fmt.Printf("%v{Id: %v, Name: %v,Type: %v, TypeDescriptor:%v}\n",
				root.GetSelector().Name(),
				root.GetId(),
				root.GetName(),
				root.GetType().Definition(),
				root.GetType().Descriptor().String(),
			)
			if len(children) == 0 {
				fmt.Printf("\t- No children\n")
			}
			for _, child := range children {
				fmt.Printf("\t- %v{Id: %v, Name: %v, Type: %v, TypeDescriptor:%v}\n",
					child.GetSelector().Name(),
					child.GetId(),
					child.GetName(),
					child.GetType().Definition(),
					child.GetType().Descriptor().String(),
				)
			}
			println()
		}
	}
}
