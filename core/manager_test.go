package core

import (
	"fmt"
	"testing"
)

func TestManager(t *testing.T) {
	obj0 := &JsObject{}
	obj0.Set("ID", "0")
	obj1 := &JsObject{}
	obj1.Set("ID", "1")
	obj2 := &JsObject{}
	obj2.Set("ID", "2")
	obj3 := &JsObject{}
	obj3.Set("ID", "3")
	obj4 := &JsObject{}
	obj4.Set("ID", "4")
	obj5 := &JsObject{}
	obj5.Set("ID", "5")

	mgr := NewJsManager()
	var err error
	err = mgr.Add(obj0)
	if err != nil {
		panic(err)
	}
	err = mgr.Add(obj1, "0")
	if err != nil {
		panic(err)
	}
	err = mgr.Add(obj2, "0")
	if err != nil {
		panic(err)
	}
	err = mgr.Add(obj3, "1", "2")
	if err != nil {
		panic(err)
	}
	err = mgr.Add(obj4, "2", "3")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Manager: %v\n\n", mgr)
}
