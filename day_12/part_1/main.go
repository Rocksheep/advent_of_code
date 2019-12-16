package main

import "fmt"

type vector struct {
	x, y, z int
}

type moon struct {
	position vector
	velocity vector
}

func (v vector) sum() int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func main() {
	moons := []moon{
		moon{vector{16, -8, 13}, vector{}},
		moon{vector{4, 10, 10}, vector{}},
		moon{vector{17, -5, 6}, vector{}},
		moon{vector{13, -3, 0}, vector{}},
	}

	simulationSteps := 1000

	for step := 0; step < simulationSteps; step++ {
		for i := 0; i < len(moons); i++ {
			moonA := &moons[i]
			for j := i + 1; j < len(moons); j++ {
				moonB := &moons[j]

				if moonA.position.x < moonB.position.x {
					moonA.velocity.x++
					moonB.velocity.x--
				} else if moonA.position.x > moonB.position.x {
					moonA.velocity.x--
					moonB.velocity.x++
				}
				if moonA.position.y < moonB.position.y {
					moonA.velocity.y++
					moonB.velocity.y--
				} else if moonA.position.y > moonB.position.y {
					moonA.velocity.y--
					moonB.velocity.y++
				}
				if moonA.position.z < moonB.position.z {
					moonA.velocity.z++
					moonB.velocity.z--
				} else if moonA.position.z > moonB.position.z {
					moonA.velocity.z--
					moonB.velocity.z++
				}
			}
		}

		for i := 0; i < len(moons); i++ {
			moons[i].position.x += moons[i].velocity.x
			moons[i].position.y += moons[i].velocity.y
			moons[i].position.z += moons[i].velocity.z
		}
	}

	energy := 0
	for _, moon := range moons {
		energy += moon.position.sum() * moon.velocity.sum()
	}
	fmt.Println(moons, energy)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
