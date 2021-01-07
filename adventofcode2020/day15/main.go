package main

import "fmt"

func main() {
	fmt.Println(GetResult([]int{7, 14, 0, 17, 11, 1, 2}, 2020))
	fmt.Println(GetResult([]int{7, 14, 0, 17, 11, 1, 2}, 30000000))
}

func GetResult(startingNumbers []int, num int) int {

	lastNum := 0
	mem := map[int][]int{}
	for t := 1; t <= num; t++ {
		//fmt.Printf("Turn %d last number spoken: %d\n", t, lastNum)

		if t <= len(startingNumbers) {
			lastNum = startingNumbers[t-1]
		} else {
			if m, found := mem[lastNum]; !found || len(m) == 1 {
				lastNum = 0
			} else {
				l := len(m)
				lastNum = m[l-1] - m[l-2]
			}
		}
		//fmt.Printf("Turn %d said: %d\n", t, lastNum)
		if _, found := mem[lastNum]; found {
			mem[lastNum] = append(mem[lastNum], t)
		} else {
			mem[lastNum] = []int{t}
		}
	}
	return lastNum
}
