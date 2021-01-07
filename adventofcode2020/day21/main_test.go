package main

import "testing"

func TestIntersect(t *testing.T) {
	m1 := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	m2 := map[string]bool{
		"b": true,
		"d": true,
	}

	res := intersect(m1, m2)

	if _, found := res["a"]; found {
		t.Error("Did not expect a")
	}

	if _, found := res["b"]; !found {
		t.Error("Expected b")
	}

	if _, found := res["c"]; found {
		t.Error("Did not expect c")
	}

	if _, found := res["d"]; found {
		t.Error("Did not expect d")
	}
}
