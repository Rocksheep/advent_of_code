package main

import "fmt"

type vector struct {
	x, y, z int
}

type moon struct {
	position vector
	velocity vector
}

func main() {
	originalState := []moon{
		moon{vector{16, -8, 13}, vector{}},
		moon{vector{4, 10, 10}, vector{}},
		moon{vector{17, -5, 6}, vector{}},
		moon{vector{13, -3, 0}, vector{}},
	}
	moons := make([]moon, len(originalState))
	copy(moons, originalState)

	xSteps, ySteps, zSteps := 0, 0, 0
	for step := 1; xSteps == 0 || ySteps == 0 || zSteps == 0; step++ {
		simulateSystem(moons)

		if xSteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.position.x != originalState[i].position.x || moon.velocity.x != originalState[i].velocity.x {
					found = false
					break
				}
			}

			if found {
				fmt.Println("X found")
				xSteps = step
			}
		}
		if ySteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.position.y != originalState[i].position.y || moon.velocity.y != originalState[i].velocity.y {
					found = false
					break
				}
			}

			if found {
				fmt.Println("Y found")
				ySteps = step
			}
		}
		if zSteps == 0 {
			found := true
			for i, moon := range moons {
				if moon.position.z != originalState[i].position.z || moon.velocity.z != originalState[i].velocity.z {
					found = false
					break
				}
			}

			if found {
				fmt.Println("Z found")
				zSteps = step
			}
		}
	}

	fmt.Println(xSteps, ySteps, zSteps)
	result := lcm(xSteps, ySteps)

	fmt.Println(lcm(result, zSteps))
}

func simulateSystem(moons []moon) []moon {
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

	return moons
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
