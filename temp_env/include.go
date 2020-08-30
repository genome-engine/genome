package temp_env

func (list *ConstList) Include(f ...Filter) *ConstList {
	var newList ConstList

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
func (list *PackList) Include(f ...Filter) *PackList {
	var newList PackList

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
func (list *StructList) Include(f ...Filter) *StructList {
	var newList StructList

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
func (list *IfaceList) Include(f ...Filter) *IfaceList {
	var newList IfaceList

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
func (list *VarList) Include(f ...Filter) *VarList {
	var newList VarList

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
func (list *FuncList) Include(f ...Filter) *FuncList {
	var newList FuncList

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
