package core

import (
	"testing"
)

func TestMap(t *testing.T) {
	obj := JsObject{}
	obj.Set("a", "b")
	obj.Set("b", 1)
	obj.Set("c", 1.55)
	obj.Set("d", uint(10))
	obj.Set("e", true)
	obj.Set(1, "k")
	println(obj.Get("a").String())
	println(obj.Get("b").Int())
	println(obj.Get("c").Float())
	println(obj.Get("d").Uint())
	println(obj.Get("e").Bool())
	println(obj.Get(1).String())
	obj.Set(10, "10000")
	println(obj.Len())
}

func TestArray(t *testing.T) {
	obj := JsObject{}
	obj.Array("a", "b", 1, 2, true, "100", 50)
	obj.SetIndex(7, "abcd")
	println(obj.GetIndex(0).String())
	println(obj.GetIndex(1).String())
	println(obj.GetIndex(2).Int())
	println(obj.GetIndex(3).Int())
	println(obj.GetIndex(4).Bool())
	println(obj.GetIndex(5).String())
	println(obj.GetIndex(6).Int())
	println(obj.GetIndex(7).String())
	obj.Append("efg")
	println(obj.GetIndex(8).String())
	println(obj.Len())
}

func TestFunc(t *testing.T) {
	obj := JsObject{func(str string) string {
		println(str)
		return str + "2"
	}}
	ret := obj.Run("bjl")
	println(ret[0].String())
}
