package main

import (
	"strings"
	"testing"
)

func BenchmarkPositions(b *testing.B) {
	input := strings.Split(Input, "\n")

	for i := 0; i < b.N; i++ {
		Positions(input)
	}
}
