package main

import (
	"fmt"
)

func usingintegers() {
	//signed integers
	q := 42
	//unsigned integers
	var w uint16 = 42
	fmt.Printf("%v , %T\n", q, q)
	fmt.Printf("%v , %T\n", w, w)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	//bit operations
	fmt.Println("bit operations")
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)

	//bit shifting
	a = 8
	fmt.Println("bit shifting")
	fmt.Println(a << 3)
	fmt.Println(a >> 3)
}
func usingBoleans() {
	var n bool = true
	i := 1 == 1
	j := 1 == 2
	fmt.Printf("%v , %T\n", n, n)
	fmt.Printf("%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)

}

func usingDecimal() {
	// default is float64
	n := 3.14
	n = 13.7e72
	n = 3.1E14
	fmt.Printf("%v , %T\n", n, n)

	//nos such as
}
func usingComplex() {
	var a complex128 = 1 + 2i
	b := 2 + 5.2i
	fmt.Printf("%v , %T\n", b, b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Printf("%v , %T\n", real(a), real(a))
	fmt.Printf("%v , %T\n", imag(a), imag(a))

	var n complex128 = complex(5, 12)
	fmt.Printf("%v , %T\n", n, n)

}

func usingStrings() {
	//string reperesent a utf-8 character
	//use double quotes while declaring a string
	s := " this is a string"
	fmt.Printf("%v , %T\n", s[2], s[2])
	//strings are immutable
	//string concatenation is possible
	//string can be converted into bytes
	b := []byte(s)
	fmt.Printf("%v , %T\n", b, b)

	// rune are utf-32 characters
	//read go package api's to understand this
	r := 'a'
	fmt.Printf("%v , %T\n", r, r)

}

func main() {

	fmt.Println("using Booleans")
	usingBoleans()
	fmt.Println("using Integers")
	usingintegers()
	fmt.Println("using Decimal")
	usingDecimal()

}
