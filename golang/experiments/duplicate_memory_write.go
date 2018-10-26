package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)
	m["a"] = ""
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			m["a"] = m["a"] + "hoge"
			fmt.Println(i, len(m["a"]))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(m["a"]))
}
