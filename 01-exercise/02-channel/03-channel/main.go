package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 6) //Con el atributo de capacidad, hasta que no se llene completamente no permite la lectura

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
