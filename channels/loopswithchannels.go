package main

import (
	"fmt"
)

//decicated incomming and sending channels
func eg4() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) //indicated to the loop that the range is till this point only
		wg.Done()
	}(ch)
	wg.Wait()
}
func eg5() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) //indicated to the loop that the range is till this point only
		wg.Done()
	}(ch)
	wg.Wait()
}
