package temp_env

func (list *ConstantList) Include(f ...Filterer) ConstantList {
	var newList ConstantList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *PackageList) Include(f ...Filterer) PackageList {
	var newList PackageList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *CustomList) Include(f ...Filterer) CustomList {
	var newList CustomList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *StructureList) Include(f ...Filterer) StructureList {
	var newList StructureList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *InterfaceList) Include(f ...Filterer) InterfaceList {
	var newList InterfaceList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *UnknownList) Include(f ...Filterer) UnknownList {
	var newList UnknownList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *VariableList) Include(f ...Filterer) VariableList {
	var newList VariableList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *MethodList) Include(f ...Filterer) MethodList {
	var newList MethodList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}

func (list *FunctionList) Include(f ...Filterer) FunctionList {
	var newList FunctionList

	if len(f) == 0 {
		return *list
	}

	if len(f) == 1 {
		return list.where(f[0], true)
	}

	newList = append(newList, list.where(f[0], true)...)

	for _, filterer := range f[1:] {
		newList = newList.where(filterer, true)
	}

	return newList
}
