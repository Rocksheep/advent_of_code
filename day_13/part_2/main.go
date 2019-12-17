package main

import (
	"fmt"
	"time"

	"advent_of_code/day_13/part_2/opcode"
)

type point struct {
	x, y int64
}

func main() {
	opCode := [99999]int64{2, 380, 379, 385, 1008, 2109, 564931, 381, 1005, 381, 12, 99, 109, 2110, 1102, 0, 1, 383, 1102, 1, 0, 382, 21001, 382, 0, 1, 21002, 383, 1, 2, 21102, 1, 37, 0, 1106, 0, 578, 4, 382, 4, 383, 204, 1, 1001, 382, 1, 382, 1007, 382, 35, 381, 1005, 381, 22, 1001, 383, 1, 383, 1007, 383, 21, 381, 1005, 381, 18, 1006, 385, 69, 99, 104, -1, 104, 0, 4, 386, 3, 384, 1007, 384, 0, 381, 1005, 381, 94, 107, 0, 384, 381, 1005, 381, 108, 1105, 1, 161, 107, 1, 392, 381, 1006, 381, 161, 1101, -1, 0, 384, 1106, 0, 119, 1007, 392, 33, 381, 1006, 381, 161, 1102, 1, 1, 384, 20101, 0, 392, 1, 21102, 19, 1, 2, 21101, 0, 0, 3, 21101, 138, 0, 0, 1105, 1, 549, 1, 392, 384, 392, 21001, 392, 0, 1, 21101, 0, 19, 2, 21102, 1, 3, 3, 21101, 161, 0, 0, 1106, 0, 549, 1101, 0, 0, 384, 20001, 388, 390, 1, 20101, 0, 389, 2, 21101, 0, 180, 0, 1105, 1, 578, 1206, 1, 213, 1208, 1, 2, 381, 1006, 381, 205, 20001, 388, 390, 1, 20102, 1, 389, 2, 21101, 205, 0, 0, 1105, 1, 393, 1002, 390, -1, 390, 1102, 1, 1, 384, 21001, 388, 0, 1, 20001, 389, 391, 2, 21101, 0, 228, 0, 1106, 0, 578, 1206, 1, 261, 1208, 1, 2, 381, 1006, 381, 253, 21002, 388, 1, 1, 20001, 389, 391, 2, 21102, 1, 253, 0, 1106, 0, 393, 1002, 391, -1, 391, 1101, 1, 0, 384, 1005, 384, 161, 20001, 388, 390, 1, 20001, 389, 391, 2, 21101, 279, 0, 0, 1105, 1, 578, 1206, 1, 316, 1208, 1, 2, 381, 1006, 381, 304, 20001, 388, 390, 1, 20001, 389, 391, 2, 21102, 1, 304, 0, 1106, 0, 393, 1002, 390, -1, 390, 1002, 391, -1, 391, 1101, 1, 0, 384, 1005, 384, 161, 21002, 388, 1, 1, 21001, 389, 0, 2, 21102, 1, 0, 3, 21101, 338, 0, 0, 1106, 0, 549, 1, 388, 390, 388, 1, 389, 391, 389, 20102, 1, 388, 1, 20101, 0, 389, 2, 21102, 1, 4, 3, 21101, 0, 365, 0, 1106, 0, 549, 1007, 389, 20, 381, 1005, 381, 75, 104, -1, 104, 0, 104, 0, 99, 0, 1, 0, 0, 0, 0, 0, 0, 216, 15, 16, 1, 1, 17, 109, 3, 21202, -2, 1, 1, 21202, -1, 1, 2, 21102, 1, 0, 3, 21102, 1, 414, 0, 1105, 1, 549, 22102, 1, -2, 1, 21202, -1, 1, 2, 21102, 429, 1, 0, 1106, 0, 601, 2102, 1, 1, 435, 1, 386, 0, 386, 104, -1, 104, 0, 4, 386, 1001, 387, -1, 387, 1005, 387, 451, 99, 109, -3, 2106, 0, 0, 109, 8, 22202, -7, -6, -3, 22201, -3, -5, -3, 21202, -4, 64, -2, 2207, -3, -2, 381, 1005, 381, 492, 21202, -2, -1, -1, 22201, -3, -1, -3, 2207, -3, -2, 381, 1006, 381, 481, 21202, -4, 8, -2, 2207, -3, -2, 381, 1005, 381, 518, 21202, -2, -1, -1, 22201, -3, -1, -3, 2207, -3, -2, 381, 1006, 381, 507, 2207, -3, -4, 381, 1005, 381, 540, 21202, -4, -1, -1, 22201, -3, -1, -3, 2207, -3, -4, 381, 1006, 381, 529, 21201, -3, 0, -7, 109, -8, 2106, 0, 0, 109, 4, 1202, -2, 35, 566, 201, -3, 566, 566, 101, 639, 566, 566, 1201, -1, 0, 0, 204, -3, 204, -2, 204, -1, 109, -4, 2106, 0, 0, 109, 3, 1202, -1, 35, 594, 201, -2, 594, 594, 101, 639, 594, 594, 20102, 1, 0, -2, 109, -3, 2106, 0, 0, 109, 3, 22102, 21, -2, 1, 22201, 1, -1, 1, 21102, 373, 1, 2, 21102, 96, 1, 3, 21101, 735, 0, 4, 21102, 630, 1, 0, 1106, 0, 456, 21201, 1, 1374, -2, 109, -3, 2106, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 2, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 2, 0, 2, 2, 0, 0, 0, 2, 2, 0, 1, 1, 0, 0, 2, 2, 0, 0, 2, 2, 2, 2, 0, 2, 2, 2, 2, 0, 2, 0, 2, 2, 2, 2, 2, 0, 0, 2, 0, 0, 2, 2, 0, 0, 0, 1, 1, 0, 2, 0, 2, 2, 2, 0, 0, 0, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 2, 0, 0, 0, 1, 1, 0, 0, 2, 0, 2, 2, 2, 2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 2, 2, 0, 0, 2, 0, 1, 1, 0, 0, 2, 0, 2, 2, 0, 0, 0, 0, 2, 2, 2, 2, 0, 2, 0, 0, 0, 2, 0, 2, 2, 0, 2, 0, 2, 2, 0, 0, 2, 2, 0, 1, 1, 0, 2, 2, 0, 2, 2, 2, 0, 0, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 0, 2, 2, 0, 2, 0, 0, 2, 0, 2, 2, 2, 0, 1, 1, 0, 2, 2, 0, 2, 0, 0, 0, 0, 0, 0, 2, 2, 2, 0, 0, 2, 2, 2, 2, 0, 2, 2, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 1, 1, 0, 2, 0, 0, 0, 2, 2, 2, 0, 0, 0, 0, 0, 2, 2, 0, 2, 0, 2, 0, 0, 2, 2, 2, 2, 0, 2, 2, 0, 0, 0, 0, 0, 1, 1, 0, 0, 2, 0, 2, 2, 0, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 2, 0, 2, 0, 0, 2, 2, 0, 0, 1, 1, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 2, 0, 2, 0, 0, 0, 2, 0, 2, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 2, 2, 2, 0, 1, 1, 0, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 0, 2, 0, 0, 2, 0, 2, 0, 2, 0, 2, 2, 2, 0, 0, 2, 0, 0, 0, 0, 2, 0, 1, 1, 0, 0, 0, 2, 0, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 0, 2, 0, 0, 0, 0, 2, 2, 2, 0, 2, 0, 0, 0, 0, 0, 1, 1, 0, 0, 2, 2, 0, 2, 0, 0, 0, 2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 2, 0, 0, 2, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 47, 10, 65, 41, 19, 85, 49, 30, 6, 77, 1, 64, 75, 29, 81, 69, 6, 16, 97, 47, 19, 46, 54, 91, 41, 16, 16, 52, 35, 72, 84, 60, 84, 52, 54, 43, 5, 24, 10, 87, 18, 10, 53, 41, 77, 61, 12, 40, 71, 1, 6, 16, 35, 59, 37, 74, 32, 11, 59, 58, 19, 11, 35, 62, 36, 48, 48, 39, 56, 5, 56, 60, 34, 92, 84, 21, 63, 37, 95, 34, 29, 67, 30, 63, 12, 62, 69, 59, 3, 53, 82, 87, 21, 90, 85, 19, 21, 58, 97, 55, 46, 78, 44, 82, 67, 37, 52, 47, 22, 62, 43, 71, 28, 84, 8, 38, 38, 95, 57, 51, 20, 96, 47, 10, 23, 56, 6, 68, 40, 48, 52, 54, 23, 70, 77, 97, 78, 87, 41, 62, 33, 2, 96, 56, 32, 65, 71, 9, 60, 69, 39, 60, 51, 95, 59, 44, 19, 33, 52, 1, 54, 9, 42, 46, 53, 30, 17, 48, 90, 79, 69, 38, 74, 46, 72, 39, 64, 59, 29, 36, 40, 13, 81, 5, 27, 98, 46, 26, 3, 48, 75, 90, 40, 52, 26, 47, 61, 7, 16, 2, 27, 10, 78, 98, 76, 30, 95, 60, 64, 5, 95, 94, 74, 79, 7, 51, 11, 84, 10, 89, 78, 35, 84, 2, 65, 74, 95, 26, 61, 27, 5, 13, 70, 65, 92, 32, 58, 53, 32, 1, 9, 15, 61, 29, 16, 72, 46, 60, 37, 66, 55, 86, 47, 72, 57, 21, 23, 30, 5, 85, 12, 22, 79, 98, 28, 87, 94, 10, 15, 82, 40, 9, 27, 40, 79, 53, 47, 33, 61, 54, 10, 42, 75, 22, 27, 51, 96, 68, 16, 98, 3, 56, 12, 81, 18, 61, 86, 33, 96, 26, 63, 87, 74, 34, 23, 22, 12, 55, 26, 62, 87, 7, 23, 14, 26, 63, 71, 98, 50, 59, 67, 34, 78, 78, 63, 34, 57, 5, 17, 29, 23, 27, 27, 17, 8, 61, 89, 81, 19, 9, 36, 70, 28, 25, 31, 20, 41, 41, 21, 30, 69, 97, 7, 38, 97, 96, 60, 90, 41, 63, 64, 74, 81, 43, 71, 65, 47, 37, 76, 37, 6, 13, 17, 82, 49, 55, 67, 7, 87, 69, 58, 63, 30, 75, 54, 41, 6, 78, 68, 37, 49, 29, 12, 77, 85, 96, 17, 36, 60, 19, 18, 19, 74, 61, 76, 83, 70, 11, 5, 66, 72, 4, 32, 6, 45, 38, 43, 22, 32, 17, 84, 2, 24, 36, 54, 7, 77, 14, 53, 57, 96, 46, 5, 14, 3, 90, 12, 95, 10, 13, 65, 78, 19, 29, 54, 91, 57, 24, 51, 69, 94, 59, 87, 46, 47, 10, 51, 51, 89, 31, 78, 21, 65, 48, 5, 52, 62, 88, 18, 20, 56, 89, 4, 23, 87, 52, 83, 8, 68, 15, 3, 80, 21, 56, 75, 45, 83, 24, 12, 30, 45, 52, 71, 4, 32, 33, 48, 8, 44, 13, 76, 61, 57, 58, 83, 97, 54, 5, 17, 64, 13, 27, 69, 91, 47, 37, 70, 4, 78, 85, 43, 82, 30, 76, 21, 11, 48, 85, 14, 79, 1, 96, 68, 46, 64, 39, 59, 37, 69, 53, 30, 49, 36, 48, 14, 37, 97, 41, 44, 90, 10, 62, 53, 62, 88, 75, 33, 31, 33, 96, 96, 82, 37, 59, 17, 88, 69, 41, 40, 19, 23, 53, 27, 3, 3, 92, 79, 82, 97, 57, 16, 23, 75, 44, 46, 7, 61, 22, 45, 28, 3, 30, 23, 2, 41, 8, 17, 27, 41, 72, 88, 48, 2, 14, 53, 49, 3, 85, 31, 23, 11, 87, 18, 15, 69, 10, 26, 74, 19, 1, 39, 98, 44, 60, 73, 72, 61, 62, 60, 84, 14, 55, 93, 65, 49, 85, 75, 88, 54, 35, 95, 91, 15, 5, 24, 45, 97, 12, 4, 46, 8, 74, 60, 70, 2, 3, 8, 68, 57, 28, 14, 11, 75, 7, 57, 42, 6, 68, 13, 35, 63, 10, 70, 53, 42, 50, 46, 34, 89, 29, 19, 83, 53, 51, 55, 63, 62, 82, 97, 28, 78, 87, 57, 13, 20, 61, 98, 19, 53, 43, 71, 20, 73, 6, 91, 81, 73, 57, 83, 50, 75, 67, 18, 95, 66, 58, 7, 14, 39, 54, 53, 35, 85, 38, 51, 91, 35, 55, 82, 47, 69, 32, 60, 88, 46, 564931}
	processor := opcode.New(opCode)

	outputCounter := 0

	var tmpX int64
	var tmpY int64
	var tmpID int64
	var score int64
	previousBallPosition := point{-1, -1}
	playerPosition := point{-1, -1}
	tiles := map[point]int64{}
	for processor.Stopped == false {
		var dX int64
		if playerPosition.x > previousBallPosition.x {
			dX = -1
		} else if playerPosition.x < previousBallPosition.x {
			dX = 1
		} else {
			dX = 0
		}

		processor.SetInput([]int64{dX})
		processor.Step()

		if processor.Waiting == true {
			output := processor.Output

			if outputCounter == 0 {
				tmpX = output
			} else if outputCounter == 1 {
				tmpY = output
			} else {
				tmpID = output

				if tmpX == -1 && tmpY == 0 {
					score = output
				} else {
					p := point{tmpX, tmpY}
					if tmpID == 4 {
						previousBallPosition = p
					} else if tmpID == 3 {
						playerPosition = p
					}
					tiles[p] = tmpID
				}

				if playerPosition.x != -1 {
					printTiles(tiles)
				}
			}

			outputCounter = (outputCounter + 1) % 3
			processor.Continue()
		}
	}
	fmt.Println("Score", score)
}

func printTiles(tiles map[point]int64) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")
	maxX := int64(0)
	maxY := int64(0)

	output := ""

	for p := range tiles {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := int64(0); y <= maxY; y++ {
		for x := int64(0); x <= maxX; x++ {
			switch (tiles[point{x, y}]) {
			case 0:
				output += " "
				break
			case 1:
				output += "#"
				break
			case 2:
				output += "X"
				break
			case 3:
				output += "-"
				break
			case 4:
				output += "o"
				break
			}
		}
		output += "\n"
	}

	fmt.Print(output)
	time.Sleep(10 * time.Millisecond)
}