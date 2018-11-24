package thirdparty

import (
	"log"
)

type A struct {
	Pub string
	pri string
}

func (a A) Print() {
	log.Println(a.Pub, a.pri)
}

func (a A) Func() {
	log.Println("function called")
}

func NewA() A {
	a := A{"Public", "Private"}
	return a
}
