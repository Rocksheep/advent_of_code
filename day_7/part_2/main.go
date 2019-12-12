package main

import (
	"advent_of_code/day_7/part_2/opcode"
	"fmt"
)

// Operation ...
type Operation struct {
	id          int
	paramsCount int
}

func main() {
	opCode := [...]int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 39, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 2, 9, 9, 102, 3, 9, 9, 1001, 9, 2, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 102, 4, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}
	phaseSettings := []int{9, 7, 8, 5, 6}

	highestValue := 0

	permutations(phaseSettings, func(setting []int) {
		var amps []*opcode.Processor
		for i := 0; i < 5; i++ {
			copy := []int{}
			for _, code := range opCode {
				copy = append(copy, code)
			}
			amps = append(amps, opcode.New(copy))
		}
		for index, phaseSettings := range phaseSettings {
			amps[index].Input = []int{phaseSettings}
		}

		ampOutput := 0
		stopped := 0
		for stopped < len(amps) {
			for _, amp := range amps {
				if amp.Stopped == true {
					stopped++
					continue
				}
				amp.AddInput(ampOutput)
				for amp.Waiting == false {
					amp.Step()
				}
				ampOutput = amp.Output
			}
		}
		if ampOutput > highestValue {
			highestValue = ampOutput
			fmt.Println("Final output:", ampOutput)
		}
	})
}

func permutations(phaseSettings []int, f func([]int)) {
	getPermutation(phaseSettings, f, 0)
}

func getPermutation(phaseSettings []int, f func([]int), i int) {
	if i > len(phaseSettings) {
		f(phaseSettings)
		return
	}
	getPermutation(phaseSettings, f, i+1)

	for j := i + 1; j < len(phaseSettings); j++ {
		phaseSettings[i], phaseSettings[j] = phaseSettings[j], phaseSettings[i]
		getPermutation(phaseSettings, f, i+1)
		phaseSettings[i], phaseSettings[j] = phaseSettings[j], phaseSettings[i]
	}
}
