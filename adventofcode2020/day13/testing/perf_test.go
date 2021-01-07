package main

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkPart2(b *testing.B) {
	input := `41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,659,x,x,x,x,x,x,x,23,x,x,x,x,13,x,x,x,x,x,19,x,x,x,x,x,x,x,x,x,29,x,937,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17`
	busses := map[int]int{}

	for i, t := range strings.Split(input, ",") {
		val, err := strconv.Atoi(t)
		if err != nil {
			continue
		}
		busses[i] = val
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findTimestamp(busses)
	}
}
