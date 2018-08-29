package main

import (
	"fmt"
	"strings"
)

// (a * b) + c -> ((a b *) c +)
//
//     +
//    / \
//   *   c
//  / \
// a   b
//
// [a, b, *, c, +]  > Binary Heap

// a + b * c -> (a + (b * c))
//
//   +
//  / \
// a  *
//   / \
//  b   c

var (
	s string = "a + b * c"
)

func tokenize(s string) []string {
	return strings.Split(s, " ")
}

func main() {
	for 
	fmt.Println(tokenize(s))
}
