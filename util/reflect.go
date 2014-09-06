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
