package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("unix time now: %d\n", t.Unix())

	duration := 2 * time.Minute
	fmt.Printf("%d minute ago: %d\n", int(duration.Minutes()), t.Add(duration).Unix())
}
