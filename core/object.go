package core

import (
	"reflect"

	"github.com/GarX/JsObject/util"
)

type JsObject struct {
	value interface{}
}

func (this *JsObject) Value() interface{} {
	return this.value
}

func pack(value interface{}) *JsObject {
	if util.IsTypeEqual(value, &JsObject{}) {
		return value.(*JsObject)
	} else {
		return &JsObject{value}
	}
}

func (this *JsObject) Set(key interface{}, value interface{}) {
	if !util.IsMap(this.value) {
		this.value = make(map[interface{}]*JsObject)
	}
	(this.value.(map[interface{}]*JsObject))[key] = pack(value)
}

func (this *JsObject) Get(key interface{}) *JsObject {
	if !util.IsMap(this.value) {
		return nil
	}
	return (this.value.(map[interface{}]*JsObject))[key]
}

func (this *JsObject) Array(elements ...interface{}) {
	array := make([]*JsObject, len(elements))
	for i := 0; i < len(elements); i++ {
		array[i] = pack(elements[i])
	}
	this.value = array
}

func (this *JsObject) SetByIndex(index int, value interface{}) {
	if !util.IsArray(this.value) {
		array := make([]*JsObject, index+1)
		array[index] = pack(value)
		this.value = array
		return
	} else {
		v := this.value.([]*JsObject)
		if l := len(v); l <= index {
			sliceToAppend := make([]*JsObject, index-l+1)
			this.value = append(v, sliceToAppend...)
			v = this.value.([]*JsObject)
		}
		v[index] = pack(value)
	}
}

func (this *JsObject) GetByIndex(index int) *JsObject {
	if !util.IsArray(this.value) {
		return nil
	}
	v := this.value.([]*JsObject)
	if l := len(v); l <= index {
		return nil
	}
	return v[index]
}

func (this *JsObject) Len() int {
	if util.IsArray(this.value) {
		v := this.value.([]*JsObject)
		return len(v)
	}
	if util.IsMap(this.value) {
		v := this.value.(map[interface{}]*JsObject)
		return len(v)
	}
	return 0
}

func (this *JsObject) Append(value interface{}) {
	if !util.IsArray(this.value) {
		return
	}
	v := this.value.([]*JsObject)
	v = append(v, pack(value))
	this.value = v
}

// Run() calls the function.
// If the value cannot be called, it reutrns a nil slice.
func (this *JsObject) Run(args ...interface{}) (ret []*JsObject) {
	if !util.IsFunc(this.value) {
		return
	}
	ret = make([]*JsObject, 0)
	refValue := reflect.ValueOf(this.value)
	refArgs := make([]reflect.Value, 0)
	for i := 0; i < len(args); i++ {
		refarg := reflect.ValueOf(args[i])
		refArgs = append(refArgs, refarg)
	}
	refReturn := refValue.Call(refArgs)
	for i := 0; i < len(refReturn); i++ {
		ret = append(ret, &JsObject{refReturn[i].Interface()})
	}
	return
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
