package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func panicing() {
	// in go we don' have exceptions as a lot of cases considered normal in go language are considered exceptions in other language
	fmt.Println("panic outputs start ")
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	fmt.Println("start")
	// panic("Something bad happended")
	fmt.Println("end")

	//panic happens after the defer statemnets a very important point to note
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	// panic("Something bad happended")
	fmt.Println("end")

}
func recovering() {
	fmt.Println("revoring start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//if we want to check the err and panic again in case of some kind of error do it here
			//panic(err)
		}
	}()
	panic("Something bad happened recover")

	// this will panic then it will recover from the panic and log the error message to the console
	//end is not executed as the panic point is where the current function stops execution
	// fmt.Println("end")
}

func main() {

	//defer works in a way that it stores all the defer funci\tion statements and executes them all after the funtion completes execution but before   the control flow returns
	// They are storwd in slack
	fmt.Println("reverse order")
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	//defer eg 2
	//allows us to follow a pattern open a resource ,check for the error , close a resource but still continue to use it inside the function
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //closed at the end of function execution
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	//another anamolous behaviour of defer keyword
	a := "start"
	defer fmt.Println(a)
	a = "end"
	//the output will be start
	panicing()
}
