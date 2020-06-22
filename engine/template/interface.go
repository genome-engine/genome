package template

import "github.com/genome-engine/genome/engine/collection"

type ITemplate interface {
	SetCollection(collection collection.Collection)
	GetSource() string
}
