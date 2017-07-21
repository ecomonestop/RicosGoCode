package main

import (
	"fmt"
)

func main() {
	//going to create a slice
	//[]<data type> is the syntax for slice type
	//slice is backed by an array under the hood
	//second arg below is the length we want for our slice
	//third arg is the capacity of the underline array that backs our slice
	//below shows example of declaring a slice with the make function
	myCourses := make([]string, 5, 10)
	fmt.Println("legnth of my slice is ", len(myCourses), "capacity of the underline array is ", cap(myCourses))

	//shorthand slice for declare and init
	myOtherCourses := []string{"course1", "course2"}
	fmt.Println(myOtherCourses)

	//accessing values of a slice
	numberSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(numberSlice[1])

	//changing value of slice
	numberSlice[3] = 0

	//slicing a slice
	//Remeber - SLICES ARE PASSED BY REFERENCE.  THIS MEANS CHANGING A VALUE IN SLICEOFNUMBERSLICE
	//EFFECTS THE ORIGINAL SLICE I.E. NUMBERSLICE BC THEY SHARE THE SAME UNDERLINE ARRAY
	sliceOfNumberSlice := numberSlice[0:2]
	fmt.Println(sliceOfNumberSlice)

	//Appending to a slice and show its resizable
	numberSlice2 := make([]int, 5, 10)
	for i := 0; i < 41; i++ {
		numberSlice2 = append(numberSlice2, i)
		fmt.Println("The capacity of numberSlice is ", cap(numberSlice2))
	}

	//append a slice to a slice
	numberSlice = append(numberSlice, numberSlice2...)
	fmt.Println(numberSlice)

}
