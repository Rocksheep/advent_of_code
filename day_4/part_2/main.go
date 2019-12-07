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
	incrementalMatch := true
	doubleMatch := false
	groupLengths := []int{}
	copiesInRow := 1

	for i := 1; i < len(value); i++ {
		curr, _ := strconv.Atoi(string(value[i]))
		prev, _ := strconv.Atoi(string(value[i-1]))
		if curr < prev {
			incrementalMatch = false
		}
		if curr == prev {
			copiesInRow++
		} else {
			groupLengths = append(groupLengths, copiesInRow)
			copiesInRow = 1
		}
	}
	groupLengths = append(groupLengths, copiesInRow)

	for _, value := range groupLengths {
		if value == 2 {
			doubleMatch = true
		}
	}

	return incrementalMatch && doubleMatch
}
