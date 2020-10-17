package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(" for loop using one var")
	}
	//using 2 vars
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println("using 2 vars")
	}

	//incorrect syntax
	//	for i := 0, j:= 0; i < 5; i= i+1, j = j+1 {

	j := 1
	for ; j < 5; j++ {
		fmt.Println("initialization outside")
	}

	// go doesn't has a do while loop
	//while loop
	for j < 10 {
		fmt.Println(" while")
		j++
	}

	// do while loop
	for {
		fmt.Println("do while loop")
		j++
		if j > 13 {
			break
		}
	}

	//continue statement same as other languages

	//using Loop labels
Loop:
	for i := 1; i < 5; i++ {
		for k := 1; k <= 3; k++ {
			fmt.Println(i * k)
			if i*k >= 3 {
				break Loop // breaks out of the outer loop
			}
		}
	}

	//for each loop using a slice and arrays and maps
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

}
