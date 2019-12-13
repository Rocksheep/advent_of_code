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
	pixelsOnLayer := width * height

	image := []int{}

	for i := 0; i < pixelsOnLayer; i++ {
		image = append(image, 2)
	}
	for i, char := range input {
		layerIndex := i % pixelsOnLayer
		value := int(char - '0')

		if image[layerIndex] == 2 {
			image[layerIndex] = value
		}
	}

	for i := 0; i < pixelsOnLayer; i += width {
		for j := 0; j < width; j++ {
			if image[i+j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
