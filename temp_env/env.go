package temp_env

import (
	"fmt"
	c "github.com/genome-engine/genome/engine/collection"
)

const (
	FileStart = "#file-start:"
	FileEnd   = "#file-end:"

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

	Collection c.Collection
	fields     []string

	logs     bool
	count    int
	filename string
}

func New(collection c.Collection, logs bool) *Env {
	collection.Linking()
	collection.ChangeQualifier("TemplateEnvironment")

	e := &Env{Collection: collection, fields: fields, logs: logs}
	e.collect()

	return e
}

func (e *Env) log(info string, args ...interface{}) {
	if !e.logs {
		return
	}
	e.count++
	fmt.Printf("\t\t\t%d.[TempEnv] %v.\n", e.count, fmt.Sprintf(info, args...))
}

func (e *Env) StartFile(filename string) string {
	e.filename = filename
	return fmt.Sprintf("%v%v\n", FileStart, filename)
}
func (e *Env) EndFile() string { return fmt.Sprintf("%v%v\n", FileEnd, e.filename) }
