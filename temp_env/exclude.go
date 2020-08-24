package temp_env

func (list *PackageList) Exclude(f ...Filter) *PackageList {
	var newList PackageList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *ConstantList) Exclude(f ...Filter) *ConstantList {
	var newList ConstantList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *VariableList) Exclude(f ...Filter) *VariableList {
	var newList VariableList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *UnknownList) Exclude(f ...Filter) *UnknownList {
	var newList UnknownList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *StructureList) Exclude(f ...Filter) *StructureList {
	var newList StructureList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *InterfaceList) Exclude(f ...Filter) *InterfaceList {
	var newList InterfaceList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *MethodList) Exclude(f ...Filter) *MethodList {
	var newList MethodList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *CustomList) Exclude(f ...Filter) *CustomList {
	var newList CustomList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
func (list *FunctionList) Exclude(f ...Filter) *FunctionList {
	var newList FunctionList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		newList = append(newList, *list.where(f[0], false)...)
		return &newList
	}

	newList = append(newList, *list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, false)
	}

	return &newList
}
