package types

import "strings"

type ChanMode int

const (
	RWMode ChanMode = iota
	Read
	Write
)

func chanMode(chandef string) ChanMode {
	switch {
	case strings.HasPrefix(chandef, "<-"):
		return Read
	case strings.HasSuffix(chandef, "<-"):
		return Write
	default:
		return RWMode
	}
}
