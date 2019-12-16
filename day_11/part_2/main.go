package main

import (
	"fmt"

	"advent_of_code/day_11/part_2/opcode"
)

type point struct {
	x, y int
}

func main() {
	opCode := [99999]int64{3, 8, 1005, 8, 319, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 28, 2, 1105, 12, 10, 1006, 0, 12, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 58, 2, 107, 7, 10, 1006, 0, 38, 2, 1008, 3, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 90, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 112, 1006, 0, 65, 1, 1103, 1, 10, 1006, 0, 91, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 144, 1006, 0, 32, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 169, 1, 109, 12, 10, 1006, 0, 96, 1006, 0, 5, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 201, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 223, 1, 4, 9, 10, 2, 8, 5, 10, 1, 3, 4, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 257, 1, 1, 9, 10, 1006, 0, 87, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 287, 2, 1105, 20, 10, 1, 1006, 3, 10, 1, 3, 4, 10, 101, 1, 9, 9, 1007, 9, 1002, 10, 1005, 10, 15, 99, 109, 641, 104, 0, 104, 1, 21102, 1, 932972962600, 1, 21101, 0, 336, 0, 1106, 0, 440, 21101, 838483681940, 0, 1, 21101, 0, 347, 0, 1106, 0, 440, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 3375393987, 0, 1, 21101, 394, 0, 0, 1105, 1, 440, 21102, 46174071847, 1, 1, 21102, 1, 405, 0, 1106, 0, 440, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 988648461076, 0, 1, 21101, 428, 0, 0, 1106, 0, 440, 21101, 0, 709580452200, 1, 21101, 439, 0, 0, 1105, 1, 440, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21102, 1, 471, 3, 21102, 461, 1, 0, 1106, 0, 504, 109, -2, 2106, 0, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 466, 467, 482, 4, 0, 1001, 466, 1, 466, 108, 4, 466, 10, 1006, 10, 498, 1102, 0, 1, 466, 109, -2, 2105, 1, 0, 0, 109, 4, 1202, -1, 1, 503, 1207, -3, 0, 10, 1006, 10, 521, 21102, 1, 0, -3, 22102, 1, -3, 1, 21201, -2, 0, 2, 21101, 0, 1, 3, 21102, 540, 1, 0, 1106, 0, 545, 109, -4, 2106, 0, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 568, 2207, -4, -2, 10, 1006, 10, 568, 22101, 0, -4, -4, 1105, 1, 636, 22102, 1, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 1, 587, 0, 1105, 1, 545, 22101, 0, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 606, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 628, 21201, -1, 0, 1, 21101, 0, 628, 0, 106, 0, 503, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2106, 0, 0}
	processor := opcode.New(opCode)

	mappedCoordinates := map[point]int64{}
	directions := []point{point{0, -1}, point{1, 0}, point{0, 1}, point{-1, 0}}
	robotDirection := 0
	currentPosition := point{0, 0}

	processor.AddInput(1)
	outputCounter := 0
	for processor.Stopped == false {
		processor.Step()

		if processor.Waiting == true {
			output := processor.Output
			if outputCounter == 0 {
				mappedCoordinates[currentPosition] = output
			} else {
				mod := 1
				if output == 0 {
					mod = -1
				}

				if robotDirection+mod < 0 {
					robotDirection = len(directions) - 1
				} else {
					robotDirection = (robotDirection + mod) % len(directions)
				}

				currentPosition.x += directions[robotDirection].x
				currentPosition.y += directions[robotDirection].y
				if value, ok := mappedCoordinates[currentPosition]; ok {
					processor.AddInput(value)
				} else {
					processor.AddInput(0)
				}
			}
			outputCounter = (outputCounter + 1) % 2
			processor.Continue()
		}
	}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for coordinate := range mappedCoordinates {
		if coordinate.x < minX {
			minX = coordinate.x
		}
		if coordinate.x > maxX {
			maxX = coordinate.x
		}
		if coordinate.y < minY {
			minY = coordinate.y
		}
		if coordinate.y > maxY {
			maxY = coordinate.y
		}
	}

	dX := maxX - minX
	dY := maxY - minY
	for x := 0; x <= dX; x++ {
		for y := 0; y <= dY; y++ {
			coordinate := point{x + minX, y + minY}
			if val, ok := mappedCoordinates[coordinate]; ok {
				if val == 1 {
					fmt.Print("#")
					continue
				}
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Println("Min and max X", minX, maxX)
	fmt.Println("Min and max Y", minY, maxY)
}
