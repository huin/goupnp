package typedesc

import "reflect"

type TypeDesc struct {
	GoType reflect.Type
}

type TypeMap map[string]TypeDesc

func (tm TypeMap) Clone() TypeMap {
	r := make(TypeMap, len(tm))
	for k, v := range tm {
		r[k] = v
	}
	return r
}
