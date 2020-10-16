package main

import "fmt"

func main() {
	if true {
		fmt.Println("The test is true")
	}
	phno := make(map[string]int)

	phno = map[string]int{
		"ujjwal":  6387406090,
		"ujjwal2": 9208012340,
		"ujjwal3": 7054110777,
	}

	//initializer ; test
	if pop, ok := phno["ujjwal"]; ok {
		fmt.Println(pop)
	}

	//operators > ,< , == , <= , >= , !=
	// || , &&

	guess := 30
	if guess >= 100 {
		fmt.Println("greater than 100")
	} else if guess >= 200 {
		fmt.Println("less than 100")
	} else {
		fmt.Println("go to hell")

	}

	choice := 5
	switch choice {
	case 1, 5, 10:
		fmt.Println("1,5,10")
	case 2, 4, 6:
		fmt.Println("2,4,6")
	default:
		fmt.Println("none")
	}

	switch choice2 := choice + 1; choice2 {
	case 1, 5, 10:
		fmt.Println("1,5,10")
	case 2, 4, 6:
		fmt.Println("2,4,6")
	default:
		fmt.Println("none")
	}

	//overlapping allowed in this syntax
	switch {
	case choice <= 10:
		fmt.Println("1,5,10")
	case choice <= 20:
		fmt.Println("2,4,6")
	default:
		fmt.Println("none")
	}

	// the default behaviour of switch id not to fall through yet we can make it to fall through using the keyword
	switch {
	case choice <= 10:
		fmt.Println("1,5,10")
		fallthrough
	case choice <= 20:
		fmt.Println("2,4,6")
	default:
		fmt.Println("none")
	}

	//another use of switch statement
	var i interface{} = 1 //interface can take any type of data
	switch i.(type) {
	case int:
		fmt.Println("it is a integer")
		break
		fmt.Println("don't execute this line")
	case float64:
		fmt.Println("it is a float")
	case string:
		fmt.Println("it is a string")
	default:
		fmt.Println("it is a another type")

	}

}
