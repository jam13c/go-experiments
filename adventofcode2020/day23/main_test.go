package main

import "testing"

func testTakeBasic(t *testing.T) {
	c := newCircle([]int{1, 2, 3, 4, 5})
	v := c.take()

	if v != [3]int{2, 3, 4} {
		t.Error("Expected 2,3,4")
	}
	if !containsAll(c.cups, []int{1, 5}) {
		t.Error("Expecting 1,5")
	}
}

func testTakeWrap(t *testing.T) {
	c := newCircle([]int{1, 2, 3, 4, 5})
	c.currentIdx = 3
	v := c.take()

	if v != [3]int{4, 5, 1} {
		t.Error("Expected 2,3,4")
	}
	if !containsAll(c.cups, []int{2, 3}) {
		t.Error("Expecting 2,3")
	}
}

func TestMinMax(t *testing.T) {
	c := newCircle([]int{1, 2, 3, 4, 5})

	if c.minCup != 1 {
		t.Error("Expected 1 for min")
	}
	if c.maxCup != 5 {
		t.Error("Expected 5 for max")
	}
}

func containsAll(cups []int, find []int) bool {
	for _, f := range find {
		found := false
		for _, c := range cups {
			if f == c {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
