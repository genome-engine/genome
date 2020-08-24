package temp_env

import "github.com/genome-engine/genome/engine/units"

func (e *Env) exist(name string, unit units.Unit) bool {
	switch name {
	case packages:
		for _, field := range *e.packs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case customs:
		for _, field := range *e.customs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case interfaces:
		for _, field := range *e.ifaces {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case imports:
		for _, field := range *e.imports {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case structures:
		for _, field := range *e.structs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case methods:
		for _, field := range *e.methods {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case functions:
		for _, field := range *e.funcs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case constants:
		for _, field := range *e.consts {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case variables:
		for _, field := range *e.vars {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case unknowns:
		for _, field := range *e.unknowns {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	}

	return false
}
