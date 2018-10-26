package main

import (
	"fmt"
	"sync"
)

func main() {
	var share string
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			share = share + "hoge"
			fmt.Println(i, len(share))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(share))
}
