package parser

import (
	"github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/parser/visitors"
)

type Config struct {
	InspectImplements    bool
	ImplementsCollection collection.Collector
	GeneralCollection    collection.Collector
	Modes                []visitors.VisitMode
	Path                 string
}
