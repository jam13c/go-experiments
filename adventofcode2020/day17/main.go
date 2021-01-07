package main

import (
	"fmt"
	"strings"
)

func main() {
	uni3 := NewUniverse(RawInput, 3)
	fmt.Println(uni3.Run(6))

	uni4 := NewUniverse(RawInput, 4)
	fmt.Println(uni4.Run(6))
}

type Universe struct {
	grid map[Vector]bool
}

func NewUniverse(input string, dimensions int) Universe {
	uni := Universe{
		grid: make(map[Vector]bool),
	}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char != '#' {
				continue
			}

			if dimensions == 3 {
				uni.grid[Vector3{x, y, 0}] = true
			} else if dimensions == 4 {
				uni.grid[Vector4{x, y, 0, 0}] = true
			} else {
				panic("Unknown dimension")
			}
		}
	}
	return uni
}

func (uni *Universe) RunCycle() {
	search := make(map[Vector]bool)

	for k, v := range uni.grid {
		search[k] = v
		for _, n := range k.Nearby() {
			if _, found := search[n]; !found {
				search[n] = false
			}
		}
	}

	newGrid := make(map[Vector]bool)
	for v, active := range search {
		activeNeighbours := 0
		for _, n := range v.Nearby() {
			if search[n] {
				activeNeighbours++
			}
		}

		if active && (activeNeighbours == 2 || activeNeighbours == 3) {
			newGrid[v] = true
		} else if !active && activeNeighbours == 3 {
			newGrid[v] = true
		}
	}
	uni.grid = newGrid
}

func (uni *Universe) countActive() int {
	c := 0
	for _, active := range uni.grid {
		if active {
			c++
		}
	}
	return c
}

func (uni *Universe) Run(c int) int {
	for i := 0; i < c; i++ {
		uni.RunCycle()
	}
	return uni.countActive()
}

type Vector3 struct {
	X, Y, Z int
}

type Vector4 struct {
	X, Y, Z, W int
}

type Vector interface {
	Nearby() []Vector
}

func (v *Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

func (v Vector3) Nearby() []Vector {
	ret, i := make([]Vector, 26), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x == 0 && y == 0 && z == 0) {
					ret[i] = v.Add(Vector3{X: x, Y: y, Z: z})
					i++
				}
			}
		}
	}

	return ret
}

func (v *Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
		W: v.W + v2.W,
	}
}

func (v Vector4) Nearby() []Vector {
	ret, i := make([]Vector, 80), 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						ret[i] = v.Add(Vector4{X: x, Y: y, Z: z, W: w})
						i++
					}
				}
			}
		}
	}

	return ret
}

const DemoInput = `.#.
..#
###`

const RawInput = `######.#
#.###.#.
###.....
#.####..
##.#.###
.######.
###.####
######.#`
