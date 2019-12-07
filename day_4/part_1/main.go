package main

import (
	"fmt"
	"strconv"
)

func main() {
	min := 138307
	max := 654504

	patternMatches := 0
	for i := min; i < max; i++ {
		if matchPattern(i) {
			patternMatches++
		}
	}

	fmt.Println("Matches: ", patternMatches)
}

func matchPattern(input int) bool {
	value := strconv.Itoa(input)
	doubleMatch := false

	for i := 1; i < len(value); i++ {
		curr, _ := strconv.Atoi(string(value[i]))
		prev, _ := strconv.Atoi(string(value[i-1]))
		if curr < prev {
			return false
		}
		if curr == prev {
			doubleMatch = true
		}
	}

	return doubleMatch
}
