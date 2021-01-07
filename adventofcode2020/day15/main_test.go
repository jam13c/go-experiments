package main

import "testing"

func TestGetResult2020(t *testing.T) {
	tables := []struct {
		n []int
		e int
	}{
		{[]int{0, 3, 6}, 436},
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}

	for _, table := range tables {
		res := GetResult(table.n, 2020)
		if res != table.e {
			t.Errorf("Result of %v was incorrect, got : %d, expected: %d", table.n, res, table.e)
		}
	}
}

func TestGetResult30000000(t *testing.T) {
	tables := []struct {
		n []int
		e int
	}{
		{[]int{0, 3, 6}, 175594},
		{[]int{1, 3, 2}, 2578},
		{[]int{2, 1, 3}, 3544142},
		{[]int{1, 2, 3}, 261214},
		{[]int{2, 3, 1}, 6895259},
		{[]int{3, 2, 1}, 18},
		{[]int{3, 1, 2}, 362},
	}

	for _, table := range tables {
		res := GetResult(table.n, 30000000)
		if res != table.e {
			t.Errorf("Result of %v was incorrect, got : %d, expected: %d", table.n, res, table.e)
		}
	}
}
