package main

import (
	"fmt"
	"reflect"
)

//Doctor struct stores the details of each doctor
type Doctor struct {
	number     int
	actorName  string `required max:"100"`
	companions []string
}

//embedding another struct inside a struct

//Physician is another struct thatembeds a Doctor struct
type Physician struct {
	Doctor
	fee     float32
	address string
}

func embeddingStructs() {
	b := Physician{}
	b.number = 6387
	b.actorName = "Emu"
	b.companions = []string{"ujjwl"}
	b.fee = 400
	b.address = "hell"
	fmt.Println("using struct ", b)
	c := Physician{
		Doctor:  Doctor{6387, "Emu", []string{"ujjwl"}},
		fee:     400,
		address: "hell",
	}

	c.Doctor.number = 333
}

func usingStructs() {
	//1st way of using a struct
	aDoctor := Doctor{
		number:    3,
		actorName: "ujjwal",
		companions: []string{
			"ujjwal1", "ujjwal2",
		},
	}
	fmt.Println("using struct ", aDoctor.companions)
	//2nd way of using the struct
	bDoctor := Doctor{
		3,
		"ujjwal",
		[]string{
			"ujjwal1", "ujjwal2",
		},
	}

	fmt.Println("using struct ", bDoctor)

	//condensed structs
	cDoctor := struct{ name string }{name: "ujjwal agarwal"}
	fmt.Println("using condensed struct ", cDoctor)

	//structs are not references they are independent
	dDoctor := cDoctor
	dDoctor.name = "mudit agarwal"
	fmt.Println(cDoctor)
	fmt.Println(dDoctor)

	eDoctor := &cDoctor
	dDoctor.name = "eeeee agarwal"
	fmt.Println(cDoctor)
	fmt.Println(eDoctor)

	//getting tags in print line using the reflect package

	t := reflect.TypeOf(Doctor{})
	field, _ := t.FieldByName("actorName")
	fmt.Println(field.Tag)
}

func main() {
	statePopulation := map[string]int{
		"ujjwal":  6387406090,
		"ujjwal2": 9208012340,
		"ujjwal3": 7054110777,
	}
	fmt.Println(statePopulation)
	//slices cannot be the key to a map
	//buy arrays can

	//usig the make function
	phno := make(map[string]int)
	phno = map[string]int{
		"ujjwal":  6387406090,
		"ujjwal2": 9208012340,
		"ujjwal3": 7054110777,
	}
	fmt.Println(phno)
	phno["hello1"] = 54444

	fmt.Println(phno["hello1"])
	fmt.Println(phno)

	delete(phno, "hello1")
	fmt.Println(phno)

	//using , ok format , which gets the boolean value depending on the availablity 0f the value for that key
	pop, ok := phno["ujjwal"]
	fmt.Println(pop, ok)

	//len function also works on maps
	fmt.Printf("lenght of the map is %v", len(phno))

	fmt.Println("using structs")
	usingStructs()
}
