package types

import (
	"regexp"
	"strings"
)

func isFunc(def string) bool  { return strings.HasPrefix(def, "func ") }
func isChan(def string) bool  { return regexp.MustCompile("^(chan |<-chan |chan<- )").MatchString(def) }
func isArray(def string) bool { return regexp.MustCompile("^\\[[0-9]+]").MatchString(def) }
func isSlice(def string) bool { return strings.HasPrefix(def, "[]") }
func isMap(def string) bool   { return strings.HasPrefix(def, "map") }
func isImported(def string) bool {
	var notAnotherType = !isFunc(def) && !isChan(def) && !isArray(def) && !isSlice(def) && !isMap(def)
	return strings.Contains(def, ".") && notAnotherType
}
func isBuiltin(def string) bool {
	builtins := []string{
		"bool", "uint", "uint8", "uint16", "uint32", "uint64",
		"int", "int8", "int16", "int32", "int64", "float32", "float64",
		"complex64", "complex128", "string", "uintptr", "byte", "rune", "error",
	}

	for _, typ := range builtins {
		if def == typ {
			return true
		}
	}

	return false
}

func IdentifyDescription(def string) TypeDescriptor {
	switch {
	case isSlice(def):
		return Slice
	case isBuiltin(def):
		return Builtin
	case isArray(def):
		return Array
	case isMap(def):
		return Map
	case isChan(def):
		return Chan
	case isFunc(def):
		return Func
	case isImported(def):
		return Imported
	case def == EmptyIface:
		return EmptyInterface
	case def == EmptyStructure:
		return EmptyStruct
	default:
		return Unknown
	}
}
