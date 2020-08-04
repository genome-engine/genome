package temp_env

import (
	c "github.com/genome-engine/genome/engine/collection"
)

const (
	methods    = "Methods"
	variables  = "Vars"
	structures = "Structs"
	unknowns   = "Unknowns"
	packages   = "Packs"
	constants  = "Consts"
	customs    = "Customs"
	imports    = "Imports"
	interfaces = "Ifaces"
	functions  = "Funcs"
)

var fields = []string{
	methods,
	variables,
	structures,
	unknowns,
	packages,
	constants,
	imports,
	interfaces,
	functions,
}

type Env struct {
	Unknowns UnknownList
	Methods  MethodList
	Customs  CustomList
	Structs  StructureList
	Imports  ImportList
	Consts   ConstantList
	Ifaces   InterfaceList
	Packs    PackageList
	Funcs    FunctionList
	Vars     VariableList

	Collection c.Collector
	fields     []string
}

func New(collection c.Collector) *Env {
	collection.Linking()

	e := &Env{Collection: collection, fields: fields}
	e.collect()

	return e
}
