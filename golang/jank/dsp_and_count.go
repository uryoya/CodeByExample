package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	conn_count = 0
)

func bidHandler(w http.ResponseWriter, r *http.Request) {
	conn_count++
	fmt.Fprintf(w, "{\"url\": \"http://www.example.com/image/999\", \"price\": 90}")
}

func main() {
	go func() {
		for {
			fmt.Printf("conn: %d\r", conn_count)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	http.HandleFunc("/bid", bidHandler)
	http.ListenAndServe(":8888", nil)
}
