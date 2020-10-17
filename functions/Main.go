package main

import (
	"fmt"
)

//pass by value
func message(msg string) {
	fmt.Println(msg)
}

//when we have multiple params with the same data type
func message2(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}

//pointers passing in the function
//pass by pointer is a more efficient method when it comes to performance
func messageReference(msg1, msg2 *string) {
	fmt.Println(*msg1, *msg2)
}

//variable no of parameters
func sum(values ...int) int {
	// simply means put all the values passed in side a slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v

	}
	return result
}

// we can also return a pointer to the varibale
// in other languages this is not  preffered operations as once the function completes execution it's local executon stack is destroyed and all the local variables as weel
// so we are logically returning address than has been destroyed but the case is not the same with go lang as then
// what go does is to contect the variable location from the stack memory to heap memory

func sum2(values ...int) *int {
	// simply means put all the values passed in side a slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v

	}
	return &result
}

// we can have a pre-declared name for the variable being declared too
func sum3(values ...int) (result int) {
	// simply means put all the values passed in side a slice
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v

	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide a number by 0")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

//method on struct pass by value
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

//method on struct pass by pointer
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
}

func main() {
	// basic syntax

	//parameters
	//return values
	//anonymous functions
	// functions as types
	// methods
	message("hello")
	message2("hello ", "second")

	greeting := "hello "
	greeting2 := "using refernce"
	messageReference(&greeting, &greeting2)
	result := sum(1, 2, 3, 4)
	fmt.Println("the sum is ", result)
	result2 := sum2(1, 2, 3, 4, 5)
	fmt.Println("the sum2 is ", *result2)

	// error handling in functions
	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//anonymous function
	func() {
		fmt.Println("anamolous function")
	}()

	// using anamolous functions inside for loops
	// prefer this syntax so that even if you are running async code the function execution becomes syncronous

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}
	// we can also use functions as variables
	f := func() {
		fmt.Println("function as varibale")
	}
	f()
	var f1 func() = func() {
		fmt.Println("function as varibale 2")
	}
	f1()
	var f2 func(a, b int) (int, float64)
	f2 = func(a, b int) (int, float64) {
		return a, 0.0
	}
	aaaa, bbbb := f2(2, 3)
	aaaa++
	bbbb++
}
