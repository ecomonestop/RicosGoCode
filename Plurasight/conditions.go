package main

import (
	"fmt"
	"os"
)

func main() {
	//Below 'rico, nico := 1,5;' is optional and is called simple statement.
	//simple statement makes these variables rico and nico scope set to only the if block
	//and is a convient way to init variables, etc for your if statement
	if rico, nico := 1, 5; rico < nico {
		fmt.Println("Rico is number 1")
	}

	//regular switch
	switch "Rico is cool" {
	case "Rico is not cool":
		fmt.Println("Isnt true")
	case "Rico is cool":
		fmt.Println("is so true")
	default:
		fmt.Println("no matches found")
	}

	//switch fall through
	//if a case is true and has a fallthrough keyword, then it executes the remaining cases below it.
	//not including the default
	switch "Rico is cool" {
	case "Rico is cool":
		fmt.Println("This is so true")
		fallthrough
	case "Rico is very cool":
		fmt.Println("even truer")
	default:
		fmt.Println("no matches found")

	}

	//switches can check many values at once
	switch "Rico is cool" {
	case "Rico is kinda cool", "Rico is so cool", "Rico is cool":
		fmt.Println("checked a few cases and see im still cool")
	default:
		fmt.Println("no matches found")
	}

	//if statements used in error hanlding to check if type error is nil or not
	//type error is usually returned as the last type returned when calling a function
	// e.g. function signature func a(a string) (x string, y string, err error)
	// notice that error is the last value returned.  its standard way
	_, err := os.Open("c:\\somefolderthatdoesntexist\\")
	if err != nil {
		fmt.Println("do some error handling")
	}
}
