package temp_env

import (
	"github.com/genome-engine/genome/engine/units"
)

//UnknownWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) UnknownWP(selector, name string) UnknownList {
	var entities UnknownList

	for _, entity := range e.Unknowns {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//UnknownWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) UnknownWC(selector, name string) UnknownList {
	var entities UnknownList

	for _, entity := range e.Unknowns {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//MethodListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) MethodWP(selector, name string) MethodList {
	var entities MethodList

	for _, entity := range e.Methods {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//MethodListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) MethodWC(selector, name string) MethodList {
	var entities MethodList

	for _, entity := range e.Methods {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//CustomListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) CustomWP(selector, name string) CustomList {
	var entities CustomList

	for _, entity := range e.Customs {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//CustomListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) CustomWC(selector, name string) CustomList {
	var entities CustomList

	for _, entity := range e.Customs {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//StructListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) StructWP(selector, name string) StructureList {
	var entities StructureList

	for _, entity := range e.Structs {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//StructListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) StructWC(selector, name string) StructureList {
	var entities StructureList

	for _, entity := range e.Structs {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//ImportListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) ImportWP(selector, name string) ImportList {
	var entities ImportList

	for _, entity := range e.Imports {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//ImportListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) ImportWC(selector, name string) ImportList {
	var entities ImportList

	for _, entity := range e.Imports {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//ConstListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) ConstWP(selector, name string) ConstantList {
	var entities ConstantList

	for _, entity := range e.Consts {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//ConstListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) ConstWC(selector, name string) *ConstantList {
	var entities ConstantList

	for _, entity := range e.Consts {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return &entities
}

//IfaceListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) IfaceWP(selector, name string) InterfaceList {
	var entities InterfaceList

	for _, entity := range e.Ifaces {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//IfaceListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) IfaceWC(selector, name string) InterfaceList {
	var entities InterfaceList

	for _, entity := range e.Ifaces {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//PackListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) PackWP(selector, name string) PackageList {
	var entities PackageList

	for _, entity := range e.Packs {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//PackListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) PackWC(selector, name string) PackageList {
	var entities PackageList

	for _, entity := range e.Packs {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//FuncListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) FuncWP(selector, name string) FunctionList {
	var entities FunctionList

	for _, entity := range e.Funcs {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//FuncListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) FuncWC(selector, name string) FunctionList {
	var entities FunctionList

	for _, entity := range e.Funcs {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//VarListWithParent:
//Returns a list of entities that have parent entities with a given selector and name
func (e *Env) VarWP(selector, name string) VariableList {
	var entities VariableList

	for _, entity := range e.Vars {
		parents, err := e.Collection.SearchParents(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, parent := range parents {
			if parent.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}

//VarListWithChild:
//Returns a list of entities that have children entities with a given selector and name
func (e *Env) VarWC(selector, name string) VariableList {
	var entities VariableList

	for _, entity := range e.Vars {
		children, err := e.Collection.SearchChildren(&entity, units.ToSelector(selector))
		if err != nil {
			continue
		}

		for _, child := range children {
			if child.GetName() == name {
				entities = append(entities, entity)
			}
		}
	}

	return entities
}
