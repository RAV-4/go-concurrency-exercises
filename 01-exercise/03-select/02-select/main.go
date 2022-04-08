package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(4 * time.Second)
		ch <- "one"
	}()

	// implement timeout for recv on channel ch

	select {
	case m := <-ch:
		fmt.Println(m)
	case <-time.After(3 * time.Second): //Con esto controlamos que si la operacion dura mas la finaliza con un timeout
		fmt.Println("Timeout")
	}
}
