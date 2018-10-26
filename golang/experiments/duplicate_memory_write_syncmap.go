package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}

	m.Store("a", "")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		mux.Lock()
		go func() {
			buf, _ := m.Load("a")
			m.Store("a", buf.(string)+"hoge")
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	s, _ := m.Load("a")
	fmt.Println(len(s.(string)))
}
