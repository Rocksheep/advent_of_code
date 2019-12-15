package main

import (
	"fmt"
	"io/ioutil"
)

type point struct {
	x, y int
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
			asteroids = append(asteroids, point{x, y})
		}
		x++
	}

	bestAsteroid := asteroids[0]
	numberOfScannedAsteroids := 0
	for index, asteroid := range asteroids {
		hits := getNumberOfScannedAsteroids(index, asteroids)
		numHits := len(hits)
		if numHits > numberOfScannedAsteroids {
			numberOfScannedAsteroids = numHits
			bestAsteroid = asteroid
		}
	}

	fmt.Println("Best location:", bestAsteroid.x, bestAsteroid.y)
	fmt.Println("Hits:", numberOfScannedAsteroids)
}

func getNumberOfScannedAsteroids(index int, asteroids []point) []point {
	station := asteroids[index]
	hits := []point{}

	for i, asteroid := range asteroids {
		if i == index {
			continue
		}

		dX := station.x - asteroid.x
		dY := station.y - asteroid.y

		divisor := abs(getGreatestCommonDivisor(dX, dY))
		p := point{dX / divisor, dY / divisor}

		alreadyExists := false
		for _, hit := range hits {
			if hit.x == p.x && hit.y == p.y {
				alreadyExists = true
				break
			}
		}
		if alreadyExists == false {
			hits = append(hits, p)
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
