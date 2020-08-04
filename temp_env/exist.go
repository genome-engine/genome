package temp_env

import "github.com/genome-engine/genome/engine/units"

func (e *Env) exist(name string, unit units.Unit) bool {
	switch name {
	case packages:
		for _, field := range e.Packs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case customs:
		for _, field := range e.Customs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case interfaces:
		for _, field := range e.Ifaces {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case imports:
		for _, field := range e.Imports {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case structures:
		for _, field := range e.Structs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case methods:
		for _, field := range e.Methods {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case functions:
		for _, field := range e.Funcs {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case constants:
		for _, field := range e.Consts {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case variables:
		for _, field := range e.Vars {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	case unknowns:
		for _, field := range e.Unknowns {
			if field.GetId() == unit.GetId() {
				return true
			}
		}
	}

	return false
}
