package temp_env

import "github.com/genome-engine/genome/engine/units"

func (e *Env) collect() {
	for _, unit := range e.Collection.Search() {
		switch u := unit.(type) {
		case *units.Variable:
			if !e.exist(variables, unit) {
				*e.vars = append(*e.vars, *u)
			}
		case *units.Constant:
			if !e.exist(constants, unit) {
				*e.consts = append(*e.consts, *u)
			}
		case *units.Function:
			if !e.exist(functions, unit) {
				*e.funcs = append(*e.funcs, *u)
			}
		case *units.Package:
			if !e.exist(packages, unit) {
				*e.packs = append(*e.packs, *u)
			}
		case *units.Import:
			if !e.exist(imports, unit) {
				*e.imports = append(*e.imports, *u)
			}
		case *units.Structure:
			if !e.exist(structures, unit) {
				*e.structs = append(*e.structs, *u)
			}
		case *units.Interface:
			if !e.exist(interfaces, unit) {
				*e.ifaces = append(*e.ifaces, *u)
			}
		case *units.Method:
			if !e.exist(methods, unit) {
				*e.methods = append(*e.methods, *u)
			}
		case *units.Unknown:
			if !e.exist(unknowns, unit) {
				*e.unknowns = append(*e.unknowns, *u)
			}
		case *units.Custom:
			if !e.exist(customs, unit) {
				*e.customs = append(*e.customs, *u)
			}
		}
	}
	e.log("Total imports: %d", len(*e.imports))
	e.log("Total packages: %d", len(*e.packs))
	e.log("Total structs: %d", len(*e.structs))
	e.log("Total interfaces: %d", len(*e.ifaces))
	e.log("Total customs: %d", len(*e.customs))
	e.log("Total functions: %d", len(*e.funcs))
	e.log("Total methods: %d", len(*e.methods))
	e.log("Total variables: %d", len(*e.vars))
	e.log("Total unknowns: %d", len(*e.unknowns))
}
