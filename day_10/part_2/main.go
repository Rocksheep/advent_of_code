package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type point struct {
	x, y, distance int
}

func main() {
	asteroids := []point{}
	y := 0
	x := 0
	data, _ := ioutil.ReadFile("./input")
	for _, char := range data {
		if char == '\n' {
			y++
			x = 0
			continue
		}
		if char == '#' {
			asteroids = append(asteroids, point{x, y, 0})
		}
		x++
	}

	numberOfScannedAsteroids := 0
	var asteroidMap map[point][]point
	for index := range asteroids {
		hits := getNumberOfScannedAsteroids(index, asteroids)
		numHits := len(hits)
		if numHits > numberOfScannedAsteroids {
			numberOfScannedAsteroids = numHits
			asteroidMap = hits
		}
	}

	angles := []float64{}
	floatMap := map[float64]point{}
	for slope := range asteroidMap {
		radians := math.Atan2(float64(slope.y), float64(slope.x))
		degrees := float64(radians * (180 / math.Pi))
		degrees += 90
		if degrees < 0 {
			degrees += 360
		}
		angles = append(angles, degrees)
		floatMap[degrees] = slope
	}
	sort.Float64s(angles)

	hitCounter := 1
	for hitCounter < 200 {
		for _, angle := range angles {
			asteroids := asteroidMap[floatMap[angle]]
			fmt.Println("Shooting", asteroids[0], " shotIndex", hitCounter)
			asteroidMap[floatMap[angle]] = asteroids[1:]
			hitCounter++

			if len(asteroidMap[floatMap[angle]]) == 0 {
				delete(asteroidMap, floatMap[angle])
			}

			if hitCounter > 200 {
				break
			}
		}
	}
}

func getNumberOfScannedAsteroids(index int, asteroids []point) map[point][]point {
	station := asteroids[index]
	hits := map[point][]point{}

	for i, asteroid := range asteroids {
		if i == index {
			continue
		}

		dX := asteroid.x - station.x
		dY := asteroid.y - station.y
		asteroid.distance = abs(dX + dY)

		divisor := abs(getGreatestCommonDivisor(dX, dY))
		p := point{dX / divisor, dY / divisor, 0}

		if _, ok := hits[p]; ok {
			hits[p] = append(hits[p], asteroid)
		} else {
			hits[p] = []point{asteroid}
		}
	}
	return hits
}

func getGreatestCommonDivisor(x int, y int) int {
	if y == 0 {
		return x
	}

	rest := x % y
	return getGreatestCommonDivisor(y, rest)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
