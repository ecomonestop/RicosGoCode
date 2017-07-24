package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//var keyword is required when declaring variables at the package level
/*var (
	name, course, module = "Rico", "Go programming", 3.6
)*/

func main() {
	//:= is a short hand within functions for declaring and assigning new variables
	name := "Rico"
	course := "Go programming"
	module := 4.5
	//&module is refencing a pointer.  it gives the memory address
	ptr := &module

	fmt.Println("Helllo from", runtime.GOOS)
	fmt.Println("name is set to ", name, "The type of name is ", reflect.TypeOf(name))
	fmt.Println("modlue is set to ", module, "The type of module is ", reflect.TypeOf(module),
		//*ptr is used to dereference a pointer.  it gives the actual value at the memory address
		//which represents the pointer
		"Module memory address is ", ptr, "and ptrs value after derefencing it is ", *ptr)

	fmt.Println("\ncourses value before calling passByValue is ", course)
	passByValue(course)
	fmt.Println("\ncourses value after calling passByValue is ", course)

	fmt.Println("\ncourses value before calling passByReference is ", course)
	passByReference(&course)
	fmt.Println("\ncourses value after calling passByReference is ", course)
}

func passByReference(course *string) string {
	*course = "This is java now and forever"
	fmt.Println("\ncourse value in passByReference is ", *course)
	return *course
}

func passByValue(course string) string {
	//notice just '=' not ':=' this is because its assigning a value to an existing variable
	course = "This is a java course now"

	fmt.Println("\nWithin passByValue course is ", course)

	return course
}
