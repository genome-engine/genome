package types

import (
	"regexp"
	"strings"
)

func extract(def string) (name, residue string) {
	switch {
	case isMap(def):
		return extMap(def)
	case isArray(def):
		return extArr(def)
	case isSlice(def):
		return def, def[2:]
	case isChan(def):
		return extChan(def)
	}

	return def, def
}

func extMap(def string) (name, residue string) {
	var start, end int

	if strings.HasPrefix(def, "map") {
		start = 3
		for i, s := range def {
			if s == ']' {
				end = i + 1
				break
			}
		}

		if end == 0 {
			return "", ""
		}

		name = def[start+1 : end-1]
		residue = def[end:]

		return name, residue
	}

	return "", ""
}

func extArr(def string) (name, residue string) {
	var arrExp = regexp.MustCompile("^\\[[0-9]+]")
	var start, end int

	if arrExp.MatchString(def) {
		start = 0

		for i, s := range def {
			if s == ']' {
				end = i + 1
				break
			}
		}

		if end == 0 {
			return "", ""
		}

		name = def[start+1 : end-1]
		residue = def[end:]

		return name, residue
	}
	return "", ""
}

func extChan(def string) (name, residue string) {
	var i = regexp.MustCompile("^(chan |<-chan |chan<- )").FindIndex([]byte(def))

	if len(i) == 0 {
		return
	}

	name = def[i[0]:i[1]]
	residue = def[i[1]:]

	return name, residue
}
