package temp_env

func (list *UnknownList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
func (list *CustomList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
func (list *FuncList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
func (list *MethodList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
func (list *ConstList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
func (list *VarList) Types() []string {
	var types []string
	for _, el := range *list {
		types = append(types, el.Type)
	}
	return types
}
