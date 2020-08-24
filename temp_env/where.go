package temp_env

import (
	"reflect"
	"strings"
)

var l = strings.ToLower

func exist(a interface{}, b ...interface{}) bool {
	for _, val := range b {
		if val == a {
			return true
		}
	}
	return false
}

func (list *StructureList) where(f Filter, eq bool) *StructureList {
	var newList StructureList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *VariableList) where(f Filter, eq bool) *VariableList {
	var newList VariableList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *MethodList) where(f Filter, eq bool) *MethodList {
	var newList MethodList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *PackageList) where(f Filter, eq bool) *PackageList {
	var newList PackageList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *CustomList) where(f Filter, eq bool) *CustomList {
	var newList CustomList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *FunctionList) where(f Filter, eq bool) *FunctionList {
	var newList FunctionList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *ConstantList) where(f Filter, eq bool) *ConstantList {
	var newList ConstantList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *InterfaceList) where(f Filter, eq bool) *InterfaceList {
	var newList InterfaceList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}

func (list *UnknownList) where(f Filter, eq bool) *UnknownList {
	var newList UnknownList

	for _, entity := range *list {
		val := reflect.ValueOf(entity)
		for i := 0; i < val.NumField(); i++ {
			name := val.Type().Field(i).Name

			if name == f.k {
				var exp bool
				iface := val.FieldByName(name).Interface()

				switch eq {
				case true:
					exp = exist(iface, f.v...)
				case false:
					exp = !exist(iface, f.v...)
				}

				if exp {
					newList = append(newList, entity)
				}
			}
		}
	}

	return &newList
}
