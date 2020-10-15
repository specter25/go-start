package main

import (
	"fmt"
	"strconv"
)

//global variable declaration
//now if we want thata variable should be available outside the package then use capital variable name

var i int = 42

var (
	name   string = "Ujjwal"
	number int    = 6387
	season int    = 11
)

var (
	address string = "hell"
)

//I this is a global variable
var I int = 25

func convert() {
	//variable conversion
	var i int = 42
	fmt.Printf("%v , %T", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v , %T", j, j)

	// to convert to strings
	//if we go the normal way it will pritn the unicode character for 43 i.t. *
	//import the strconv
	var str string
	str = strconv.Itoa(i)
	fmt.Printf("%v , %T", str, str)

}

func main() {
	// local variabes
	var i int
	i = 42
	//
	var j int = 27
	//
	k := 29
	fmt.Println(i, j, k)
	fmt.Printf("%v , %T", k, k)
	convert()

}
