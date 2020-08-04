package temp_env

func (list *StructureList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *InterfaceList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *ConstantList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *VariableList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *FunctionList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *PackageList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *UnknownList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *MethodList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *CustomList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *ImportList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
