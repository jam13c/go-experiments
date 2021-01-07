package main

import (
	"testing"
)

func TestPush(t *testing.T) {

	h := hand{}
	h.push(1)
	h.push(2)

	if len(h) != 2 {
		t.Error("Expected 2 elements")
	}
}

func TestPop(t *testing.T) {
	h := hand{}
	h.push(1)
	h.push(2)

	v := h.pop()

	if v != 1 {
		t.Error("Expected to pop 1")
	}

	if len(h) != 1 {
		t.Errorf("Expected length 1 got:%d", len(h))
	}
}
