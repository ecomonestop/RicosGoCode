package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//A map that key is the String value from std input and value of how many times this string appeared in std input - can use to spot duplicates
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for text, count := range counts {
		fmt.Println("\nThis string: ", text, " appeared ", count, " times")
	}

}
