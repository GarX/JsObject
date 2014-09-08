package core

import (
	"github.com/GarX/JsObject/constant"
	"github.com/GarX/JsObject/util"
)

// JsManager is only for map JsObject
// the key "ID" is meaningful for this package,
type JsManager struct {
	objects map[string]*JsObject
	parent  map[string](map[string]*JsObject) // get parents by childID
	child   map[string](map[string]*JsObject) // get children by parentID
}

func NewJsManager() *JsManager {
	manager := &JsManager{
		map[string]*JsObject{},
		map[string](map[string]*JsObject){},
		map[string](map[string]*JsObject){},
	}
	return manager
}

func verify(obj *JsObject) error {
	if !util.IsMap(obj.Value()) {
		return constant.NotMapError
	}
	id := obj.Get("ID").String()
	if id == "" {
		return constant.EmptyIDError
	} else {
		return nil
	}
}

func (this *JsManager) Get(id string) *JsObject {
	obj, exist := this.objects[id]
	if !exist {
		return nil
	} else {
		return obj
	}
}

func (this *JsManager) AddRelation(childID string, parentID string) {
	childObject := this.Get(childID)
	parentObject := this.Get(parentID)

	// check the map is nil or not
	parentMap, exist := this.parent[childID]
	if !exist {
		parentMap = map[string]*JsObject{}
		this.parent[childID] = parentMap
	}
	childMap, exist := this.child[parentID]
	if !exist {
		childMap = map[string]*JsObject{}
		this.child[parentID] = childMap
	}

	// add relation
	parentMap[parentID] = parentObject
	childMap[childID] = childObject
}

func (this *JsManager) RemoveRelation(childID, parentID string) {
	delete(this.parent[childID], parentID)
	delete(this.child[parentID], childID)
}

func (this *JsManager) Add(obj *JsObject, parents ...string) error {
	if err := verify(obj); err != nil {
		return err
	}
	id := obj.Get("ID").String()
	this.objects[id] = obj
	for i := 0; i < len(parents); i++ {
		this.AddRelation(id, parents[i])
	}
	return nil
}
