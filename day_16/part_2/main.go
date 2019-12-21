package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	message := parseMessage(input)
	offset := offset(string(input))

	for phase := 1; phase <= 100; phase++ {
		result := make([]int, len(message))
		sliceLength := offset

		off := sliceLength - 1

		sum := 0
		for i := len(message) - 1; i >= off; i-- {
			sum += (message[i])
			result[i] = sum % 10
		}

		copy(message, result)
	}

	for _, value := range message[offset : offset+8] {
		fmt.Print(value)
	}
	fmt.Print("\n")
}

func parseMessage(message []byte) []int {
	messageLength := len(message)
	result := make([]int, messageLength*10000)
	for i := 0; i < 10000; i++ {
		start := messageLength * i
		for j, c := range message {
			result[start+j] = int(c - '0')
		}
	}

	return result
}

func offset(input string) int {
	offset, _ := strconv.Atoi(string(input[:7]))
	return offset
}
