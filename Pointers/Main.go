package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	//creatinga pointer
	//dereferencing a pointer
	//the new function
	//working with nil
	// types of internal  pointers

	a := 42
	b := a
	fmt.Println(a)
	fmt.Println(b)
	a = 25
	fmt.Println(a)
	fmt.Println(b)

	var aa int = 5
	var bb *int = &aa
	fmt.Println(aa, bb)
	fmt.Println(&aa, *bb)

	//Pointer arithmetic
	aaa := []int{1, 2, 3}
	bbb := &aaa[0]
	ccc := &aaa[1]

	fmt.Printf("%v %p %p \n", aaa, bbb, ccc)

	//we cannot do addition and subtractions etc with the pointers in go

	var ms myStruct
	ms = myStruct{foo: 42}
	fmt.Println(ms)

	//ms2 now holds he address of the instance of the struct
	var ms2 *myStruct
	ms2 = &myStruct{foo: 42}
	fmt.Println(ms2)

	//we can also use the new function to initialie the object
	var ms3 *myStruct
	fmt.Println(ms3)
	ms3 = new(myStruct)
	//this special syntax is followed due to the operator precedence ()>.>*
	(*ms3).foo = 42
	//now but go prevents us from using this type of syntax
	// this also works as go already knows that is is the pointer and thus we dn't need to dereference it
	//this is just syntax sugar
	ms3.foo = 9
	fmt.Println(ms3)
}
