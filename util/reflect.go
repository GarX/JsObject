package util

import "reflect"

func IsMap(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind().String() == "map"
}
func IsArray(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind().String() == "slice"
}

func IsFunc(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind().String() == "func"
}

func IsTypeEqual(i0 interface{}, i1 interface{}) bool {
	type0 := reflect.TypeOf(i0)
	type1 := reflect.TypeOf(i1)
	return type0 == type1
}
