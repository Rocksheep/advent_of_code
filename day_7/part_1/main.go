package main

import (
	"fmt"
	"strconv"
)

// Operation ...
type Operation struct {
	id          int
	paramsCount int
}

func main() {
	opCode := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 30, 39, 64, 81, 102, 183, 264, 345, 426, 99999, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 5, 9, 101, 2, 9, 9, 102, 3, 9, 9, 1001, 9, 2, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 102, 4, 9, 9, 1001, 9, 5, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}

	phaseSettings := []int{1, 2, 5, 4, 3}

	highestValue := 0

	permutations(phaseSettings, func(setting []int) {
		value := 0
		for _, phaseSetting := range phaseSettings {
			result := execOpCode(opCode, []int{phaseSetting, value})
			value = result
		}

		if value > highestValue {
			highestValue = value
			fmt.Println("Highest value: ", highestValue)
			fmt.Println("Phasesetting", setting)
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

func execOpCode(opCode []int, inputs []int) int {
	addOperation := Operation{1, 3}
	multiplyOperation := Operation{2, 3}
	inputOperation := Operation{3, 1}
	outputOperation := Operation{4, 1}
	jumpIfTrueOperation := Operation{5, 2}
	jumpIfFalseOperation := Operation{6, 2}
	lessThanOperation := Operation{7, 3}
	equalsOperation := Operation{8, 3}

	exit := 99
	inputCounter := 0

	var i = 0
	for i < len(opCode) {
		currentOperation := opCode[i]
		var operation int
		var bitModes uint8 = 0

		if currentOperation > 9 {
			temp := strconv.Itoa(currentOperation)
			operation, _ = strconv.Atoi(temp[len(temp)-2:])
			bits := temp[:len(temp)-2]
			result, _ := strconv.ParseUint(bits, 2, 8)
			bitModes = uint8(result)
		} else {
			operation = currentOperation
		}

		if operation == exit {
			break
		}

		switch operation {
		case addOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			paramB := getValue(bitModes>>1, opCode, i+2)
			resultPos := opCode[i+3]

			opCode[resultPos] = paramA + paramB

			i += addOperation.paramsCount + 1
			break
		case multiplyOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			paramB := getValue(bitModes>>1, opCode, i+2)
			resultPos := opCode[i+3]

			opCode[resultPos] = paramA * paramB

			i += multiplyOperation.paramsCount + 1
			break
		case inputOperation.id:
			resultPos := opCode[i+1]

			input := inputs[inputCounter]
			inputCounter++

			opCode[resultPos] = input
			i += inputOperation.paramsCount + 1
			break
		case outputOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			return paramA
			i += outputOperation.paramsCount + 1
			break
		case jumpIfTrueOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			resultPos := getValue(bitModes>>1, opCode, i+2)
			if paramA != 0 {
				i = resultPos
			} else {
				i += jumpIfTrueOperation.paramsCount + 1
			}
			break
		case jumpIfFalseOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			resultPos := getValue(bitModes>>1, opCode, i+2)

			if paramA == 0 {
				i = resultPos
			} else {
				i += jumpIfFalseOperation.paramsCount + 1
			}
			break
		case lessThanOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			paramB := getValue(bitModes>>1, opCode, i+2)
			resultPos := opCode[i+3]

			var result = 0
			if paramA < paramB {
				result = 1
			}
			opCode[resultPos] = result
			i += lessThanOperation.paramsCount + 1
			break
		case equalsOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			paramB := getValue(bitModes>>1, opCode, i+2)
			resultPos := opCode[i+3]

			var result = 0
			if paramA == paramB {
				result = 1
			}
			opCode[resultPos] = result
			i += equalsOperation.paramsCount + 1
			break
		}
	}

	return 0
}

func getValue(mode uint8, opCode []int, index int) int {
	if mode&1 == 1 {
		return opCode[index]
	}

	return opCode[opCode[index]]
}
