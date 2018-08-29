package main

import (
	"fmt"
	"sort"
)

/* Heap機能の定義 */

type Heap interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Heap) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push(h Heap, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Heap) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove(h Heap, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

func Fix(h Heap, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Heap, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Heap, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break // j1がオーバーフローしたときとかインデクスの値がおかしいときのバリデーション
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2 // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

/* IntHeapの実装 */

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h IntHeap) Print() {
	fmt.Println(h)
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	Init(h)
	Push(h, 3)
	Push(h, 10)
	Push(h, 2)
	Push(h, 45)
	Push(h, 666)
	Push(h, 0)
	Push(h, -1)
	fmt.Printf("minimum: %d\n", (*h)[0])
	h.Print()
	for h.Len() > 0 {
		fmt.Printf("%d ", Pop(h))
	}
	fmt.Println()
}
