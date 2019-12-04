package main

import "fmt"

func main() {
	originalOperations := [...]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 2, 9, 19, 23, 1, 23, 6, 27, 1, 13, 27, 31, 1, 31, 10, 35, 1, 9, 35, 39, 1, 39, 9, 43, 2, 6, 43, 47, 1, 47, 5, 51, 2, 10, 51, 55, 1, 6, 55, 59, 2, 13, 59, 63, 2, 13, 63, 67, 1, 6, 67, 71, 1, 71, 5, 75, 2, 75, 6, 79, 1, 5, 79, 83, 1, 83, 6, 87, 2, 10, 87, 91, 1, 9, 91, 95, 1, 6, 95, 99, 1, 99, 6, 103, 2, 103, 9, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 115, 6, 119, 2, 6, 119, 123, 1, 10, 123, 127, 1, 127, 5, 131, 1, 131, 2, 135, 1, 135, 5, 0, 99, 2, 0, 14, 0}
	found := false

	for noun := 0; noun <= 99 && !found; noun++ {
		for verb := 0; verb <= 99 && !found; verb++ {
			opCode := originalOperations
			opCode[1] = noun
			opCode[2] = verb

			result := execOpCode(opCode)

			if result[0] == 19690720 {
				fmt.Println(100*noun + verb)
				found = true
			}
		}
	}

}

func execOpCode(opCode [145]int) [145]int {
	add := 1
	multiply := 2
	exit := 99

	for i := 0; i < len(opCode); i += 4 {
		instruction := opCode[i]

		if instruction == exit {
			break
		}

		paramA := opCode[opCode[i+1]]
		paramB := opCode[opCode[i+2]]
		resultPos := opCode[i+3]

		if instruction == add {
			opCode[resultPos] = paramA + paramB
		} else if instruction == multiply {
			opCode[resultPos] = paramA * paramB
		}
	}

	return opCode
}
