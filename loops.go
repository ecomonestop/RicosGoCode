package main

import (
	"fmt"
)

func main() {
	//for is the only keyword for loops
	//below is a simple loop
	for timeBeforeBoom := 10; timeBeforeBoom >= 0; timeBeforeBoom-- {
		if timeBeforeBoom == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timeBeforeBoom)
		//time.Sleep(1 * time.Second)
	}

	//below is a for range loop
	courses := []string{"java", "golang", "python"}
	//range returns two values, the index and the value of the itterable type you passed it
	//we ignore the index in this case making the var name _
	//we print out the value course
	for _, course := range courses {
		fmt.Println(course)
	}

	//Will show example of continue keyword below.  However wanted to mention
	//break labels.  break lables allow you not to just break out of the current loop
	//up to the next most outer, but gives you flexiability to choose which outer loop to break too
	//below we print odd numbers using keyword continue
	for number := 10; number >= 0; number-- {
		if number%2 == 0 {
			continue
		}
		fmt.Println(number)
	}

}
