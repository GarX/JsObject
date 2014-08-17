package util

import "reflect"

func IsMap(v reflect.Value) bool {
	if v.Kind().String() == "map" {
		return true
	}
	return false
}

func IsArray(v reflect.Value) bool {
	if v.Kind().String() == "slice" {
		return true
	}
	return false
}
