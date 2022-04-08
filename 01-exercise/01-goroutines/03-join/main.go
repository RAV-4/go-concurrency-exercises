package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int
	var wg sync.WaitGroup //objeto que nos sirve para controlar la sincronia y los flujos de las gourutines

	wg.Add(1) //Con esto podemos decir que estamos proximos a inciar un flujo concurrente
	go func() {
		defer wg.Done() //Al final del flujo concurrente debemos indicar que ya terminó, para decrementar en uno
		data++
	}()

	wg.Wait() //Este espera que el contador esta en cero, lo que le indica que ya los flujos concurrentes terminaron y puede continuar
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
