package opcode

import (
	"fmt"
	"strconv"
)

// Processor An memory processor
type Processor struct {
	memory       [99999]int64
	address      int64
	inputAddress int
	relativeBase int64
	Stopped      bool
	Waiting      bool
	Input        []int64
	Output       int64
}

// New Create a new instance of a processor
func New(memory [99999]int64) *Processor {
	return &Processor{
		memory,
		0,
		0,
		0,
		false,
		false,
		[]int64{},
		0,
	}
}

// Step Steps over the command in the address
func (processor *Processor) Step() {
	currentOperation := processor.memory[processor.address]
	var operation int64
	bitModes := []uint8{0, 0, 0}

	if currentOperation > 9 {
		temp := strconv.FormatInt(currentOperation, 10)
		operation, _ = strconv.ParseInt(temp[len(temp)-2:], 10, 64)
		modes := temp[:len(temp)-2]
		for i, mode := range modes {
			convertedMode := uint8(mode - '0')
			bitModes[len(modes)-1-i] = convertedMode
		}
	} else {
		operation = currentOperation
	}

	processor.Waiting = false

	switch operation {
	case 99:
		processor.Waiting = true
		processor.Stopped = true
		break
	case 1:
		processor.addOperation(bitModes)
		break
	case 2:
		processor.multiplyOperation(bitModes)
		break
	case 3:
		processor.inputOperation(bitModes)
		break
	case 4:
		processor.outputOperation(bitModes)
		break
	case 5:
		processor.jumpIfTrueOperation(bitModes)
		break
	case 6:
		processor.jumpIfFalseOperation(bitModes)
		break
	case 7:
		processor.lessThanOperation(bitModes)
		break
	case 8:
		processor.equalsOperation(bitModes)
		break
	case 9:
		processor.relativeBaseOperation(bitModes)
		break
	}
}

func (processor *Processor) addOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	paramB := processor.getValue(bitModes[1], 2)
	processor.setValue(paramA+paramB, bitModes[2], 3)

	processor.address += 4
}

func (processor *Processor) multiplyOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	paramB := processor.getValue(bitModes[1], 2)
	processor.setValue(paramA*paramB, bitModes[2], 3)

	processor.address += 4
}

func (processor *Processor) inputOperation(bitModes []uint8) {
	input := processor.Input[processor.inputAddress]
	processor.inputAddress++
	processor.setValue(input, bitModes[0], 1)
	processor.address += 2
}

func (processor *Processor) outputOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	processor.Output = paramA
	fmt.Println("Output time", paramA)
	processor.address += 2
	processor.Waiting = true
}

func (processor *Processor) jumpIfTrueOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	resultPos := processor.getValue(bitModes[1], 2)
	if paramA != 0 {
		processor.address = resultPos
	} else {
		processor.address += 3
	}
}

func (processor *Processor) jumpIfFalseOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	resultPos := processor.getValue(bitModes[1], 2)
	if paramA == 0 {
		processor.address = resultPos
	} else {
		processor.address += 3
	}
}

func (processor *Processor) lessThanOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	paramB := processor.getValue(bitModes[1], 2)

	var result int64 = 0
	if paramA < paramB {
		result = 1
	}
	processor.setValue(result, bitModes[2], 3)

	processor.address += 4
}

func (processor *Processor) equalsOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	paramB := processor.getValue(bitModes[1], 2)

	var result int64 = 0
	if paramA == paramB {
		result = 1
	}
	processor.setValue(result, bitModes[2], 3)

	processor.address += 4
}

func (processor *Processor) relativeBaseOperation(bitModes []uint8) {
	paramA := processor.getValue(bitModes[0], 1)
	processor.relativeBase += paramA
	processor.address += 2
}

// AddInput add Input to list of inputs
func (processor *Processor) AddInput(input int64) {
	processor.Input = append(processor.Input, input)
	processor.Waiting = false
	processor.Stopped = false
}

// Setmemory sets the memory
func (processor *Processor) Setmemory(memory [99999]int64) {
	processor.memory = memory
	processor.address = 0
}

func (processor *Processor) getValue(mode uint8, increment int64) int64 {
	address := processor.address + increment
	if mode == 1 {
		return processor.memory[address]
	} else if mode == 2 {
		address = processor.relativeBase + processor.memory[address]
		return processor.memory[address]
	}

	return processor.memory[processor.memory[address]]
}

func (processor *Processor) setValue(value int64, mode uint8, increment int64) {
	address := processor.address + increment
	if mode == 0 {
		processor.memory[processor.memory[address]] = value
		return
	}
	address = processor.relativeBase + processor.memory[address]
	processor.memory[address] = value
}
