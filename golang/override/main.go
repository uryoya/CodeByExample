// Goでオーバーライド
// https://qiita.com/Jxck_/items/509459ee7be940290130
package main

import (
	"./thirdparty"
	"log"
)

type Wrap struct {
	thirdparty.A
}

func (w Wrap) Print() {
	w.Print()
	log.Println(w.Pub)
	// log.Println(w.pri) // w.pri undefined
}

func main() {
	a := thirdparty.NewA()
	a.Print()

	w := Wrap{a}
	w.Func()
	w.Print()
}
