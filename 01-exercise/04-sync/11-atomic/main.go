package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	// TODO: implement concurrency safe counter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
				//Esto obliga que la operacion se ha de forma atomica, por lo cual es bloqueada para que otra gourutine la lea
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
