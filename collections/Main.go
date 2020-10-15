package main

import (
	"fmt"
)

func arraysGo() {
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades : %v \n", grades)

	//just declare the size sufficient enough to take the no of elements given in the array
	grades1 := [...]int{97, 85, 93}
	fmt.Printf("Grades1 : %v \n", grades1)

	var students [3]string
	fmt.Printf("Students : %v \n", students)
	students[0] = "Lisa"
	fmt.Printf("Students : %v \n", students)
	fmt.Printf("Students : %v \n", len(students))

	//Syntax 1
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	//Syntax 2
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix)
	fmt.Println(identityMatrix2)

	//arrays in go are not pointers they are absolute values
	//so if a:=b then b is a separate array not pointing to a

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// to do that we ahve to tell b to point to the same data that a has

	aa := [...]int{1, 2, 3}
	bb := &aa
	bb[1] = 5
	fmt.Println(aa)
	fmt.Println(bb)

}
func main() {
	fmt.Println("arrays")
	arraysGo()

}
