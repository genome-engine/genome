package visitors

import (
	"github.com/genome-engine/genome/engine/collection"
)

func NewGeneralVisitor(path, packMainDir, src string, collector collection.Collection, modes ...VisitMode) *GeneralVisitor {
	collector.ChangeQualifier("Parsing:Visiting")
	genVis := &GeneralVisitor{
		path: path, src: src, Collection: collector, packMainDir: packMainDir, modes: modes,
	}

	return genVis
}
func NewImportVisitor(src string, collector collection.Collection) *ImportVisitor {
	return &ImportVisitor{src: src, Collector: collector}
}
func NewStructVisitor(src string, collector collection.Collection) *StructsVisitor {
	return &StructsVisitor{src: src, Collector: collector}
}
func NewInterfaceVisitor(src string, collector collection.Collection) *InterfacesVisitor {
	return &InterfacesVisitor{src: src, Collector: collector}
}
func NewFuncVisitor(src string, collector collection.Collection) *FuncsVisitor {
	return &FuncsVisitor{src: src, Collection: collector}
}
func NewCustomVisitor(src string, collector collection.Collection) *CustomsVisitor {
	return &CustomsVisitor{src: src, Collector: collector}
}
