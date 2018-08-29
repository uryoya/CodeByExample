// arrayとsliceの違い
// https://blog.golang.org/go-slices-usage-and-internals

package main

import (
	"fmt"
)

func main() {
	// array
	// arrayは配列そのもの
	a1 := [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)

	// slice
	// sliceはarrayの参照とインデクスと長さを持つ
	//
	// Nil slice
	//     +--+
	// ptr |  |
	//     +--+
	// len |  |
	//     +--+
	// cap |  |
	//     +--+
	//
	// Slice
	//        +--+
	//        |  |----------+
	//        +--+          |
	// []byte | 5|          |
	//        +--+         +--+--+--+--+--+
	//        | 5|  [5]byte|  |  |  |  |  |
	//        +--+         +--+--+--+--+--+
	s := []int{1, 2, 3}
	fmt.Printf("s: %v\n", s)
}
