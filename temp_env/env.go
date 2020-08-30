package temp_env

import (
	"fmt"
	c "github.com/genome-engine/genome/engine/collection"
	"github.com/genome-engine/genome/engine/units"
	"strings"
)

const (
	ParentFilter = "Parent"
	ChildFilter  = "Child"

	FileStart = "#file-start:"
	FileEnd   = "#file-end:"

	methods    = "methods"
	variables  = "vars"
	structures = "structs"
	unknowns   = "unknowns"
	packages   = "packs"
	constants  = "consts"
	customs    = "customs"
	imports    = "imports"
	interfaces = "ifaces"
	functions  = "funcs"
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

type env struct {
	unknowns *UnknownList
	methods  *MethodList
	customs  *CustomList
	structs  *StructList
	imports  *ImportList
	consts   *ConstList
	ifaces   *IfaceList
	packs    *PackList
	funcs    *FuncList
	vars     *VarList

	fields []string

	logs     bool
	count    int
	filename string
}

type Env struct {
	*env
	Collection c.Collection
}

func New(collection c.Collection, logs bool) *Env {
	collection.Linking()
	collection.ChangeQualifier("TemplateEnvironment")

	e := &Env{
		env: &env{
			unknowns: &UnknownList{},
			methods:  &MethodList{},
			customs:  &CustomList{},
			structs:  &StructList{},
			imports:  &ImportList{},
			consts:   &ConstList{},
			ifaces:   &IfaceList{},
			packs:    &PackList{},
			funcs:    &FuncList{},
			vars:     &VarList{},

			fields: fields,
			logs:   logs,
		},

		Collection: collection,
	}
	e.env.collect(collection)

	return e
}

func (e *env) log(info string, args ...interface{}) {
	if !e.logs {
		return
	}
	e.count++
	fmt.Printf("\t\t\t%d.[TempEnv] %v.\n", e.count, fmt.Sprintf(info, args...))
}
func (e *Env) StartFile(filename string) string {
	e.env.filename = filename
	return fmt.Sprintf("%v%v\n", FileStart, filename)
}
func (e *Env) EndFile() string { return fmt.Sprintf("%v%v\n", FileEnd, e.filename) }

func parseValue(value string) (units.Selector, string) {
	elements := strings.Split(value, ":")
	switch {
	case len(elements) == 1:
		return units.ToSelector(elements[0]), ""
	case len(elements) > 1:
		return units.ToSelector(elements[0]), elements[1]
	default:
		return nil, ""
	}
}

func (e *Env) Unknowns(f ...Filter) *UnknownList {
	if len(f) == 0 {
		return e.unknowns
	}

	var list = &UnknownList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, unknown := range *e.env.unknowns {
					if children, _ := e.Collection.SearchChildren(&unknown, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, unknown)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, unknown := range *e.env.unknowns {
					if children, _ := e.Collection.SearchParents(&unknown, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, unknown)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Methods(f ...Filter) *MethodList {
	if len(f) == 0 {
		return e.methods
	}

	var list = &MethodList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, method := range *e.env.methods {
					if children, _ := e.Collection.SearchChildren(&method, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, method)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, method := range *e.env.methods {
					if children, _ := e.Collection.SearchParents(&method, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, method)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Customs(f ...Filter) *CustomList {
	if len(f) == 0 {
		return e.customs
	}

	var list = &CustomList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, custom := range *e.env.customs {
					if children, _ := e.Collection.SearchChildren(&custom, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, custom)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, custom := range *e.env.customs {
					if children, _ := e.Collection.SearchParents(&custom, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, custom)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Structs(f ...Filter) *StructList {
	if len(f) == 0 {
		return e.structs
	}

	var list = &StructList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))
			switch filter.k {

			case ChildFilter:
				for _, structure := range *e.env.structs {
					if children, _ := e.Collection.SearchChildren(&structure, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, structure)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, structure := range *e.env.structs {
					if children, _ := e.Collection.SearchParents(&structure, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, structure)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}

	return list
}
func (e *Env) Imports(f ...Filter) *ImportList {
	if len(f) == 0 {
		return e.imports
	}

	var list = &ImportList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, imp := range *e.env.imports {
					if children, _ := e.Collection.SearchChildren(&imp, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, imp)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, imp := range *e.env.imports {
					if children, _ := e.Collection.SearchParents(&imp, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, imp)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Consts(f ...Filter) *ConstList {
	if len(f) == 0 {
		return e.consts
	}

	var list = &ConstList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, constant := range *e.env.consts {
					if children, _ := e.Collection.SearchChildren(&constant, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, constant)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, constant := range *e.env.consts {
					if children, _ := e.Collection.SearchParents(&constant, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, constant)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Ifaces(f ...Filter) *IfaceList {
	if len(f) == 0 {
		return e.ifaces
	}

	var list = &IfaceList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, iface := range *e.env.ifaces {
					if children, _ := e.Collection.SearchChildren(&iface, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, iface)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, iface := range *e.env.ifaces {
					if children, _ := e.Collection.SearchParents(&iface, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, iface)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Packs(f ...Filter) *PackList {
	if len(f) == 0 {
		return e.packs
	}

	var list = &PackList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, pack := range *e.env.packs {
					if children, _ := e.Collection.SearchChildren(&pack, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, pack)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, pack := range *e.env.packs {
					if children, _ := e.Collection.SearchParents(&pack, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, pack)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Funcs(f ...Filter) *FuncList {
	if len(f) == 0 {
		return e.funcs
	}

	var list = &FuncList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, function := range *e.env.funcs {
					if children, _ := e.Collection.SearchChildren(&function, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, function)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, function := range *e.env.funcs {
					if children, _ := e.Collection.SearchParents(&function, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, function)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
func (e *Env) Vars(f ...Filter) *VarList {
	if len(f) == 0 {
		return e.vars
	}

	var list = &VarList{}

	for _, filter := range f {
		for _, val := range filter.v {
			selector, value := parseValue(val.(string))

			switch filter.k {

			case ChildFilter:
				for _, variable := range *e.env.vars {
					if children, _ := e.Collection.SearchChildren(&variable, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, variable)
							}
						}
					}
				}
				continue

			case ParentFilter:
				for _, variable := range *e.env.vars {
					if children, _ := e.Collection.SearchParents(&variable, selector); children != nil {
						for _, child := range children {
							if child.GetName() == value {
								*list = append(*list, variable)
							}
						}
					}
				}
				continue

			default:
				continue
			}
		}
	}
	return list
}
