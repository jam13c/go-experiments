package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	busses := map[int]int{}

	for i, t := range strings.Split(strings.Split(RawInput, "\n")[1], ",") {
		val, err := strconv.Atoi(t)
		if err != nil {
			continue
		}
		busses[i] = val
	}
	fmt.Println(busses)

	fmt.Println(findTimestamp(busses))
}

func findTimestamp(buses map[int]int) int {
	t := buses[0]
	d := buses[0]

	for dt, bus := range buses {
		for {
			if (t+dt)%bus == 0 {
				break
			}
			t += d
		}
		d = lcm(d, bus)
	}

	return t
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	return a * b / gcd(a, b)
}

type bus struct {
	t int
	i int
}

const DemoInput = `939
7,13,x,x,59,x,31,19`

const RawInput = `1000104
41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,659,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,x,x,19,x,x,x,x,x,x,x,x,x,29,x,937,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17`
