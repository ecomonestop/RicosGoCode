package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//Arguments start at index 1; whereas the command starts at index 0
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("\n", s)
}
