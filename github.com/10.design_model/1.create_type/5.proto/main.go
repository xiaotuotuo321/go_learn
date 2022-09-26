package main

import "log"

type Cloneable interface {
	Clone() Cloneable
}

type ProtoTypeManager struct {
	protoTypes map[string]Cloneable
}

func NewProtoTypeManager() *ProtoTypeManager {
	return &ProtoTypeManager{
		protoTypes: make(map[string]Cloneable),
	}
}

func (p *ProtoTypeManager) Get(name string) Cloneable {
	return p.protoTypes[name]
}

func (p *ProtoTypeManager) Set(name string, protoType Cloneable) {
	p.protoTypes[name] = protoType
}

var manager *ProtoTypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func main() {
	manager = NewProtoTypeManager()
	t10 := &Type1{
		name: "Type1",
	}
	manager.Set("t1", t10)

	t1 := manager.Get("t1")
	t2 := t1.Clone()

	if t1 == t2 {
		log.Fatalf("error! get clone not working")
	}

	c := manager.Get("t1").Clone()
	t3 := c.(*Type1)

	if t3.name != "Type1" {
		log.Fatalf("error")
	}
}
