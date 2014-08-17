package core

import (
	"reflect"

	"github.com/GarX/JsObject/util"
)

type JsObject struct {
	value interface{}
}

func (this *JsObject) Value(v interface{}) {
	this.value = v
}

func (this *JsObject) Set(key interface{}, value interface{}) {
	refthis := reflect.ValueOf(this.value)
	if !util.IsMap(refthis) {
		this.value = make(map[interface{}]*JsObject)
		refthis = reflect.ValueOf(this.value)
	}
	obj := &(JsObject{value})
	refKey := reflect.ValueOf(&key).Elem()
	refValue := reflect.ValueOf(&obj).Elem()
	refthis.SetMapIndex(refKey, refValue)
}

func (this *JsObject) Get(key interface{}) *JsObject {
	refthis := reflect.ValueOf(this.value)
	if !util.IsMap(refthis) {
		return nil
	}
	refKey := reflect.ValueOf(&key).Elem()
	refValue := refthis.MapIndex(refKey)
	return refValue.Interface().(*JsObject)
}

func (this *JsObject) Array(elements ...interface{}) {
	array := make([]*JsObject, len(elements))
	for i := 0; i < len(elements); i++ {
		obj := &JsObject{elements[i]}
		array[i] = obj
	}
	this.value = array
}

func (this *JsObject) Bool() bool {
	switch this.value.(type) {
	case bool:
		return this.value.(bool)
	case *bool:
		return *this.value.(*bool)
	default:
		return false
	}
}

func (this *JsObject) String() string {
	switch this.value.(type) {
	case string:
		return this.value.(string)
	case *string:
		return *(this.value.(*string))
	default:
		return ""
	}
}

func (this *JsObject) Int() int {
	switch this.value.(type) {
	case int:
		return this.value.(int)
	case int8:
		return (int)(this.value.(int8))
	case int16:
		return (int)(this.value.(int16))
	case int32:
		return (int)(this.value.(int32))
	case int64:
		return (int)(this.value.(int64))
	case *int:
		return *this.value.(*int)
	case *int8:
		return (int)(*this.value.(*int8))
	case *int16:
		return (int)(*this.value.(*int16))
	case *int32:
		return (int)(*this.value.(*int32))
	case *int64:
		return (int)(*this.value.(*int64))
	default:
		return 0
	}
}

func (this *JsObject) Float() float64 {
	switch this.value.(type) {
	case float32:
		return (float64)(this.value.(float32))
	case float64:
		return this.value.(float64)
	case *float32:
		return (float64)(*this.value.(*float32))
	case *float64:
		return *this.value.(*float64)
	default:
		return 0.0
	}
}

func (this *JsObject) Uint() uint {
	switch this.value.(type) {
	case uint:
		return this.value.(uint)
	case uint8:
		return (uint)(this.value.(uint8))
	case uint16:
		return (uint)(this.value.(uint16))
	case uint32:
		return (uint)(this.value.(uint32))
	case uint64:
		return (uint)(this.value.(uint64))
	case *uint:
		return *this.value.(*uint)
	case *uint8:
		return (uint)(*this.value.(*uint8))
	case *uint16:
		return (uint)(*this.value.(*uint16))
	case *uint32:
		return (uint)(*this.value.(*uint32))
	case *uint64:
		return (uint)(*this.value.(*uint64))
	default:
		return 0
	}
}
