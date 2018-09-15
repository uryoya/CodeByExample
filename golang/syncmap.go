// スレッドセーフなMap
package main

import (
	"fmt"
	"sync"
)

type Bar struct {
	i int
}

type Foo struct {
	Bars []Bar
	Str  string
}

func main() {
	m := sync.Map{}

	// 値の保存
	m.Store("hoge", "fuga")

	foo := Foo{
		Bars: []Bar{{1}, {2}},
		Str:  "hello",
	}

	m.Store("foo", foo)

	fmt.Println(foo)

	// 値の取得
	if v, ok := m.Load("hoge"); ok {
		fmt.Printf("hoge: %s\n", v)
	}

	// 既存のキーが無ければ値の保存、あれば取得
	actual, loaded := m.LoadOrStore("hoge", "piyo")
	fmt.Printf("hoge: %s, loaded=%t\n", actual, loaded)

	// 全ての要素を表示
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("Key: %s, value: %s\n", k, v)
		return true
	})

	// 値の上書き
	m.Store("hoge", "hao123")

	m.Range(func(k, v interface{}) bool {
		fmt.Printf("Key: %s, value: %s\n", k, v)
		return true
	})

	// 削除
	m.Delete("hoge")

}
