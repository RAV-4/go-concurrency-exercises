package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	//var flag bool
	//var mu sync.Mutex

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			//TODO: modify so that load function gets called only once.
			//Implementacion machetera que hace que load solo se ejecute una sola vex
			/*mu.Lock()
			if !flag {
				load()
				flag = true
			}
			mu.Unlock()*/
			once.Do(load)
			// Con once restringimos la ejecucion a una sola vez
		}()
	}
	wg.Wait()
}
