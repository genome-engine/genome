package temp_env

type Filter struct {
	k string
	v []interface{}
}

func NewFilter(k string, v ...interface{}) Filter { return Filter{k: k, v: v} }
