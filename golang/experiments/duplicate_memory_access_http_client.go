package main

import (
	"net/http"
)

func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			http.Get("http://localhost:9999/")
		}()
	}
}
