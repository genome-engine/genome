package temp_env

import (
	"reflect"
	"regexp"
	"strings"
)

var l = strings.ToLower

func exist(a interface{}, b ...interface{}) bool {
	var str = func(val interface{}) bool {
		return reflect.TypeOf(val).String() == reflect.String.String()
	}

	for _, val := range b {
		if str(val) && str(a) {
			if regexp.MustCompile(a.(string)).MatchString(val.(string)) {
				return true
			}
			continue
		}
		if val == a {
			return true
		}
	}
	return false
}

func (list *StructList) where(f Filter, eq bool) *StructList {
	var newList StructList

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

func (list *VarList) where(f Filter, eq bool) *VarList {
	var newList VarList

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

func (list *PackList) where(f Filter, eq bool) *PackList {
	var newList PackList

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

func (list *FuncList) where(f Filter, eq bool) *FuncList {
	var newList FuncList

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

func (list *ConstList) where(f Filter, eq bool) *ConstList {
	var newList ConstList

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

func (list *IfaceList) where(f Filter, eq bool) *IfaceList {
	var newList IfaceList

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
