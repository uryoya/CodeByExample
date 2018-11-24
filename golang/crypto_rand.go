package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		token := make([]byte, 32)
		rand.Read(token)
		fmt.Println(token)
	}
}
