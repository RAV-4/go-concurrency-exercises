package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}

	}()

	// TODO: if there is no value on channel, do not block.
	for i := 0; i < 2; i++ {
		//con el select con default operamos con el channel solo cuando este tenga la info
		// Para no hacerlo bloqueante
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("No message received")
		}

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)
	}
}
