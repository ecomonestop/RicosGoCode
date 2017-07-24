package main

import (
	"fmt"
)

func main() {
	type person struct {
		FirstName   string
		LastName    string
		Occumpation string
	}

	//2 ways declaring a variable as type Person
	var Rico person
	Celine := new(person)

	//Declaring and init
	Rico2 := person{
		FirstName:   "Rico",
		LastName:    "Romero",
		Occumpation: "Software Engineer",
	}
	fmt.Println(Rico, Celine, Rico2)

}
