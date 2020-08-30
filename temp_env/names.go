package temp_env

func (list *UnknownList) Names() []string {
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
func (list *StructList) Names() []string {
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
func (list *IfaceList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *ConstList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *PackList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *FuncList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
func (list *VarList) Names() []string {
	var names []string
	for _, el := range *list {
		names = append(names, el.GetName())
	}
	return names
}
