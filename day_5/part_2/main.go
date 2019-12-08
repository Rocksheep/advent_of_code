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
	opCode := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 57, 23, 224, 101, -1311, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1102, 57, 67, 225, 102, 67, 150, 224, 1001, 224, -2613, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 2, 179, 213, 224, 1001, 224, -469, 224, 4, 224, 102, 8, 223, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1001, 188, 27, 224, 101, -119, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 223, 224, 223, 1, 184, 218, 224, 1001, 224, -155, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 21, 80, 224, 1001, 224, -101, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1101, 67, 39, 225, 1101, 89, 68, 225, 101, 69, 35, 224, 1001, 224, -126, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 1102, 7, 52, 225, 1102, 18, 90, 225, 1101, 65, 92, 225, 1002, 153, 78, 224, 101, -6942, 224, 224, 4, 224, 102, 8, 223, 223, 101, 6, 224, 224, 1, 223, 224, 223, 1101, 67, 83, 225, 1102, 31, 65, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1007, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 359, 1001, 223, 1, 223, 1107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 374, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 404, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 419, 1001, 223, 1, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 449, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 494, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 509, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 539, 1001, 223, 1, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 554, 101, 1, 223, 223, 108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 629, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

	execOpCode(opCode)
}

func execOpCode(opCode []int) []int {
	addOperation := Operation{1, 3}
	multiplyOperation := Operation{2, 3}
	inputOperation := Operation{3, 1}
	outputOperation := Operation{4, 1}
	jumpIfTrueOperation := Operation{5, 2}
	jumpIfFalseOperation := Operation{6, 2}
	lessThanOperation := Operation{7, 3}
	equalsOperation := Operation{8, 3}

	exit := 99

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

			fmt.Println("Enter input: ")
			var input int
			fmt.Scanf("%d", &input)

			opCode[resultPos] = input
			i += inputOperation.paramsCount + 1
			break
		case outputOperation.id:
			paramA := getValue(bitModes, opCode, i+1)
			fmt.Println(paramA)
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

	return opCode
}

func getValue(mode uint8, opCode []int, index int) int {
	if mode&1 == 1 {
		return opCode[index]
	}

	return opCode[opCode[index]]
}
