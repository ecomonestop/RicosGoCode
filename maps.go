package main

import (
	"fmt"
)

func main() {
	//Declaring Maps
	ricoMaps := make(map[string]int)
	//setting values
	ricoMaps["number of years in IT"] = 5
	fmt.Println(ricoMaps)

	//Declaring and init maps
	celineMap := map[string]int{
		"number of years in IT": 1,
	}
	fmt.Println(celineMap)

	//Itterate over maps and prove that they're unordered list
	ricosSayings := map[int]string{1: "happy wife happy life", 2: "logic doesnt always win", 3: "football is awesome", 4: "football and beer is even better"}
	for key, value := range ricosSayings {
		fmt.Println("key: ", key, "\nvalue: ", value)
	}

	//Add to a map
	ricosSayings[5] = "No when to hold them and when to fold em"
	//Update a map
	ricosSayings[1] = "Happy wife Happy life is very true"
	fmt.Println(ricosSayings)
	//delete from map
	delete(ricosSayings, 5)
	fmt.Println(ricosSayings)

}
