package collection

import (
	"fmt"
	"github.com/genome-engine/genome/engine/units"
)

type (
	Collection struct {
		logs          bool
		count         int
		qualifier     string
		unitsMap      UnitsMap
		rootTable     map[int]units.Unit
		childrenTable map[int][]units.Unit
	}

	UnitsMap map[units.Unit][]units.Unit
)

func New(qualifier string, logs bool) *Collection {
	c := &Collection{
		logs:          logs,
		qualifier:     qualifier,
		unitsMap:      map[units.Unit][]units.Unit{},
		rootTable:     map[int]units.Unit{},
		childrenTable: map[int][]units.Unit{},
	}
	c.log("Collection was crated.")
	return c
}

func (c *Collection) log(info string, args ...interface{}) {
	if !c.logs {
		return
	}
	c.count++
	fmt.Printf("\t\t%d.[Collection:%v]%v\n", c.count, c.qualifier, fmt.Sprintf(info, args...))
}

func (c *Collection) ChangeQualifier(qualifier string) {
	c.qualifier = qualifier
	c.count = 0
}
