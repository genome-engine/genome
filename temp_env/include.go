package temp_env

func (list *ConstantList) Include(f ...Filter) *ConstantList {
	var newList ConstantList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *PackageList) Include(f ...Filter) *PackageList {
	var newList PackageList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *CustomList) Include(f ...Filter) *CustomList {
	var newList CustomList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *StructureList) Include(f ...Filter) *StructureList {
	var newList StructureList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *InterfaceList) Include(f ...Filter) *InterfaceList {
	var newList InterfaceList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *UnknownList) Include(f ...Filter) *UnknownList {
	var newList UnknownList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *VariableList) Include(f ...Filter) *VariableList {
	var newList VariableList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *MethodList) Include(f ...Filter) *MethodList {
	var newList MethodList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
func (list *FunctionList) Include(f ...Filter) *FunctionList {
	var newList FunctionList

	if len(f) == 0 {
		return list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, *list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = *newList.where(filterer, true)
	}

	return &newList
}
