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
}
