package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "functional basics"
	author := "Rico"

	fmt.Println(converter(module, author))

	fmt.Println(bestLeagueFinishDetrmined(11, 2, 34, 11, 28, 16, 3))
}

func bestLeagueFinishDetrmined(finishes ...int) (bestFinish int) {
	best := finishes[0]
	for _, place := range finishes {
		if place < best {
			best = place
		}
	}
	return best
}

func converter(titleCase, allCaps string) (titleCaseStr, allCapsStr string) {
	titleCase = strings.Title(titleCase)
	allCaps = strings.ToUpper(allCaps)

	return titleCase, allCaps
}
