package main

import (
	"fmt"
	"sync"
)

//normally in languages a single operating system thread is created to run a process
//go creted light weighted abstraction of threads that can run on normal os threads rather than creating special threads for program execution
// only the scheduler has to regulate the different threads
//and these are go-runtimes

var wg = sync.WaitGroup{}
var count = 0

func main() {
	go sayHello()

	for i := 0; i < 10; i++ {
		wg.Add(2)
		// go.sayHello()
		// go.increment()
	}
	wg.Wait()
}
func sayHello() {
	fmt.Printf("Hello ")
}
