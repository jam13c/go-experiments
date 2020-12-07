package main

import (
	"regexp"
	"testing"
)

func BenchmarkTraverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		seatID(traverse("BFFFBBFRRR"))
	}
}

func BenchmarkRegex(b *testing.B) {
	re := regexp.MustCompile("B|F|R|L")
	for n := 0; n < b.N; n++ {
		seatIDRegex(re, "BFFFBBFRRR")
	}
}
