package util

import "reflect"

func IsMap(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind().String() == "map" {
		return true
	}
	return false
}

func IsArray(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind().String() == "slice" {
		return true
	}
	return false
}

func IsFunc(i interface{}) bool {
	v := reflect.ValueOf(i)
	if v.Kind().String() == "func" {
		return true
	}
	return false
}
