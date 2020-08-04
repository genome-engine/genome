package temp_env

func (list *VariableList) Exclude(f ...Filterer) VariableList {
	var newList VariableList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *FunctionList) Exclude(f ...Filterer) FunctionList {
	var newList FunctionList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *CustomList) Exclude(f ...Filterer) CustomList {
	var newList CustomList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *StructureList) Exclude(f ...Filterer) StructureList {
	var newList StructureList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *ConstantList) Exclude(f ...Filterer) ConstantList {
	var newList ConstantList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], false)
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *InterfaceList) Exclude(f ...Filterer) InterfaceList {
	var newList InterfaceList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *PackageList) Exclude(f ...Filterer) PackageList {
	var newList PackageList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *MethodList) Exclude(f ...Filterer) MethodList {
	var newList MethodList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}

func (list *UnknownList) Exclude(f ...Filterer) UnknownList {
	var newList UnknownList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		newList = append(newList, list.where(f[0], false)...)
		return newList
	}

	newList = append(newList, list.where(f[0], false)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, false)
	}

	return newList
}
