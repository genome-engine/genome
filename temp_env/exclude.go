package temp_env

func (list *PackList) Exclude(f ...Filter) *PackList {
	var newList PackList

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
func (list *ConstList) Exclude(f ...Filter) *ConstList {
	var newList ConstList

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
func (list *VarList) Exclude(f ...Filter) *VarList {
	var newList VarList

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
func (list *StructList) Exclude(f ...Filter) *StructList {
	var newList StructList

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
func (list *IfaceList) Exclude(f ...Filter) *IfaceList {
	var newList IfaceList

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
func (list *FuncList) Exclude(f ...Filter) *FuncList {
	var newList FuncList

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
