package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "03036732577212944063491565474664"
	encoding := []int{0, 1, 0, -1}

	offset, _ := strconv.Atoi(string(input[:7]))
	inputLength := len(input)
	bigInputLength := inputLength * 10000
	halfInputLength := inputLength / 2
	dX := bigInputLength - offset

	fmt.Println(bigInputLength, offset)

	fmt.Println(input)

	for phase := 0; phase < 100; phase++ {
		newInput := ""
		for round := 1; round <= inputLength; round++ {
			sum := 0
			i := offset
			if round > dX {
				start := round - dX - 1
				for i := halfInputLength + start; i < inputLength; i++ {
					sum += int(input[i] - '0')
					i++
				}
			} else {
				for i < bigInputLength {
					j := i % inputLength
					c := input[j]
					encodingIndex := ((j + 1) / round) % len(encoding)
					value := int(c - '0')
					sum += encoding[encodingIndex] * value
					i++
				}
			}
			newInput += strconv.Itoa(abs(sum % 10))
		}
		input = newInput
		fmt.Println(input)
	}
	fmt.Println(input[:8])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
