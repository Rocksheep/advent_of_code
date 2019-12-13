package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("./input")
	input := string(data)

	width := 25
	height := 6
	layerLength := width * height

	leastZeros := 999999
	var result map[rune]int

	counter := map[rune]int{}
	for i, char := range input {
		if i%layerLength == 0 {
			counter['0'] = 0
			counter['1'] = 0
			counter['2'] = 0
		}

		if char >= '0' && char <= '2' {
			counter[char]++
		}

		if i%layerLength == layerLength-1 && counter['0'] < leastZeros {
			leastZeros = counter['0']

			result = map[rune]int{}
			for k, v := range counter {
				result[k] = v
			}
		}
	}

	fmt.Println("result", result['1']*result['2'])
}
