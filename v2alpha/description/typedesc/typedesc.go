package typedesc

import "reflect"

type TypeDesc struct {
	GoType reflect.Type
}

type TypeMap map[string]TypeDesc
