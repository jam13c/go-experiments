package main

import "fmt"

func main() {
	fmt.Println(plus(vals()))
	fmt.Println(vPlus(vals()))
}

func vals() (int, int) {
	return 3, 7
}

func plus(a, b int) int {
	return a + b
}

func vPlus(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
