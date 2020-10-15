package main

import (
	"fmt"
)

const myconst2 = iota // counter creted when creating contanst
//will evauate to 0

// this is a constant block just like variable block
//but value to iota is rescricted to a constant block
//it's value again resets to 0 for
const (
	aa = iota
	bb = iota
	cc = iota
)

const (
	aa2 = iota
	bb2
	cc2
)

//_ says to the compiler that i know you will generate a value here but i don't need it any way
const (
	_ = iota
	bb3
	cc3
)

func main() {
	// if 2 constants with same name the inner block constant wins
	const myConst int = 42

	//a constant can be added to a variable but only if it has the same data type ....

	fmt.Printf("%v , %T \n", myConst, myConst)

	const a = 42
	var b int16 = 27
	fmt.Printf("%v , %T \n", a+b, a+b) // consts unlike variable sre implicitly converted if no given tyoe specified

	//package level iotta constants
	fmt.Printf("%v , %T \n", aa, aa)
	fmt.Printf("%v , %T \n", bb, bb)
	fmt.Printf("%v , %T \n", cc, cc)
	fmt.Printf("%v , %T \n", aa2, aa2)
	fmt.Printf("%v , %T \n", bb2, bb2)
	fmt.Printf("%v , %T \n", cc2, cc2)

}
