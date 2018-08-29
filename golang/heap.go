// container/heapパッケージを使って二分木を実装する
// 参考: http://d.hatena.ne.jp/taknb2nch/20140126/1390720293
package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// 型を定義して、この型に heap interface を実装
type myIntArray []int

func (h myIntArray) Len() int {
	return len(h)
}

func (h myIntArray) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myIntArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myIntArray) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myIntArray) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	m := &myIntArray{}

	rnd := rand.New(rand.NewSource(time.Now().Unix()))

	heap.Init(m)

	for i := 0; i < 10; i++ {
		v := rnd.Intn(100)
		fmt.Printf("push: %d\n", v)

		// 必ずheap経由でPush
		// NG: m.Push(v)
		heap.Push(m, v)
	}

	for _, v := range *m {
		fmt.Printf("%v ", v)
	}

	fmt.Println()

	for m.Len() > 0 {
		v1 := (*m)[0]
		v2 := heap.Pop(m)

		fmt.Printf("min: %v, pop : %v\n", v1, v2)
	}
}
