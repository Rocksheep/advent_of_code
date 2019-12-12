package opcode

import (
	"strconv"
)

// Processor An OPcode processor
type Processor struct {
	opCode       []int
	address      int
	inputAddress int
	Stopped      bool
	Waiting      bool
	Input        []int
	Output       int
}

// New Create a new instance of a processor
func New(opCode []int) *Processor {
	return &Processor{
		opCode,
		0,
		0,
		false,
		false,
		[]int{},
		0,
	}
}

// Step Steps over the command in the address
func (processor *Processor) Step() {
	currentOperation := processor.opCode[processor.address]
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
		processor.inputOperation()
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
	}
}

func (processor *Processor) addOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	paramB := processor.value(bitModes>>1, processor.address+2)
	resultAddress := processor.opCode[processor.address+3]

	processor.opCode[resultAddress] = paramA + paramB

	processor.address += 4
}

func (processor *Processor) multiplyOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	paramB := processor.value(bitModes>>1, processor.address+2)
	resultAddress := processor.opCode[processor.address+3]

	processor.opCode[resultAddress] = paramA * paramB

	processor.address += 4
}

func (processor *Processor) inputOperation() {
	// fmt.Println("Input time")
	resultAddress := processor.opCode[processor.address+1]
	input := processor.Input[processor.inputAddress]
	processor.inputAddress++

	processor.opCode[resultAddress] = input
	processor.address += 2
}

func (processor *Processor) outputOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	processor.Output = paramA
	// fmt.Println("Output time", paramA)
	processor.address += 2
	processor.Waiting = true
}

func (processor *Processor) jumpIfTrueOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	resultPos := processor.value(bitModes>>1, processor.address+2)
	if paramA != 0 {
		processor.address = resultPos
	} else {
		processor.address += 3
	}
}

func (processor *Processor) jumpIfFalseOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	resultPos := processor.value(bitModes>>1, processor.address+2)
	if paramA == 0 {
		processor.address = resultPos
	} else {
		processor.address += 3
	}
}

func (processor *Processor) lessThanOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	paramB := processor.value(bitModes>>1, processor.address+2)
	resultAddress := processor.opCode[processor.address+3]

	var result = 0
	if paramA < paramB {
		result = 1
	}
	processor.opCode[resultAddress] = result
	processor.address += 4
}

func (processor *Processor) equalsOperation(bitModes uint8) {
	paramA := processor.value(bitModes, processor.address+1)
	paramB := processor.value(bitModes>>1, processor.address+2)
	resultAddress := processor.opCode[processor.address+3]

	var result = 0
	if paramA == paramB {
		result = 1
	}
	processor.opCode[resultAddress] = result
	processor.address += 4
}

// SetInput add Input to list of inputs
func (processor *Processor) AddInput(input int) {
	processor.Input = append(processor.Input, input)
	// processor.inputAddress = 0
	// processor.address = 0
	processor.Waiting = false
	processor.Stopped = false
}

// SetOpCode sets the opcode
func (processor *Processor) SetOpCode(opCode []int) {
	processor.opCode = opCode
	processor.address = 0
}

func (processor *Processor) value(mode uint8, address int) int {
	if mode&1 == 1 {
		return processor.opCode[address]
	}

	return processor.opCode[processor.opCode[address]]
}
