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

func slicesGo() {
	//slice points to an  array
	//this is a slice if we don't put in the 3 dots

	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Students : %v \n", len(a))
	fmt.Printf("Students : %v \n", cap(a))

	//slices are naturally referenec type and refer to the same underlying data
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//other ways to create a sice
	aa := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	bb := aa[:] // slice of all the elements
	cc := aa[3:]
	dd := aa[:6]
	ee := aa[3:6]
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	//all these slices are references
	aa[5] = 42
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	//using build in make functio n to make slices
	// the thirs argument is capacity '
	// lsice has 3 elements and underlying array will have 100 elements
	aaa := make([]int, 3, 100)
	fmt.Println(aaa)
	fmt.Printf("make : %v \n", len(aaa))
	fmt.Printf("make : %v \n", cap(aaa))

	// to add an element to the slice
	bbb := []int{}
	fmt.Println(bbb)
	fmt.Printf("make bbb: %v \n", len(bbb))
	fmt.Printf("make bbb: %v \n", cap(bbb))

	bbb = append(bbb, 1)
	fmt.Println(bbb)
	fmt.Printf("append : %v \n", len(bbb))
	fmt.Printf("append : %v \n", cap(bbb))

	bbb = append(bbb, 2, 3, 4, 5, 6)
	fmt.Println(bbb)
	fmt.Printf("append multiple values  : %v \n", len(bbb))
	fmt.Printf("append multiple values: %v \n", cap(bbb))

	//using the spread opearator to concatenate 2 slices
	bbb = append(bbb, []int{7, 8, 9, 10}...)
	fmt.Println(bbb)
	fmt.Printf("append multiple values  : %v \n", len(bbb))
	fmt.Printf("append multiple values: %v \n", cap(bbb))

}

func main() {
	fmt.Println("arrays")
	arraysGo()
	fmt.Println("slice")
	slicesGo()

}
