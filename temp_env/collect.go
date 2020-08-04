package temp_env

import "github.com/genome-engine/genome/engine/units"

func (e *Env) collect() {
	for _, unit := range e.Collection.Search() {
		switch u := unit.(type) {
		case *units.Variable:
			if !e.exist(variables, unit) {
				e.Vars = append(e.Vars, *u)
			}
		case *units.Constant:
			if !e.exist(constants, unit) {
				e.Consts = append(e.Consts, *u)
			}
		case *units.Function:
			if !e.exist(functions, unit) {
				e.Funcs = append(e.Funcs, *u)
			}
		case *units.Package:
			if !e.exist(packages, unit) {
				e.Packs = append(e.Packs, *u)
			}
		case *units.Import:
			if !e.exist(imports, unit) {
				e.Imports = append(e.Imports, *u)
			}
		case *units.Structure:
			if !e.exist(structures, unit) {
				e.Structs = append(e.Structs, *u)
			}
		case *units.Interface:
			if !e.exist(interfaces, unit) {
				e.Ifaces = append(e.Ifaces, *u)
			}
		case *units.Method:
			if !e.exist(methods, unit) {
				e.Methods = append(e.Methods, *u)
			}
		case *units.Unknown:
			if !e.exist(unknowns, unit) {
				e.Unknowns = append(e.Unknowns, *u)
			}
		case *units.Custom:
			if !e.exist(customs, unit) {
				e.Customs = append(e.Customs, *u)
			}
		}
	}
}
