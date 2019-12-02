package main

import (
	"testing"
)

func TestPeekNormal(t *testing.T) {
	i := []int{1, 9, 10, 3}
	cpu := CPU{intcode: i}

	peek := cpu.Peek()
	if !SlicesEqual(peek, i[0:4]) {
		t.Fatal("unexpected read order:", peek)
	}
}

func TestPeekOverrun(t *testing.T) {
	i := []int{1, 9, 10, 3}
	cpu := CPU{intcode: i, ip: 4}

	peek := cpu.Peek()
	if !SlicesEqual(peek, []int{0, 0, 0, 0}) {
		t.Fatal("unexpected read order:", peek)
	}
}

func TestAdd(t *testing.T) {
	i := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	cpu := CPU{intcode: i}

	t.Log("IN: ", cpu.intcode)

	peek := cpu.Peek()
	if !SlicesEqual(peek, i[0:4]) {
		t.Fatal("unexpected read order:", peek)
	}

	b := cpu.DoNext()

	if !b {
		t.Errorf("DoNext returned false")
	}
	if i[3] != 70 {
		t.Errorf("add failed: expected 70, got %d\n", i[3])
	}

	if cpu.Value() != 70 {
		t.Errorf("bad value: expected=70, got=%d\n", cpu.Value())
	}

	t.Log("OUT:", cpu.intcode)
}

func TestHalt(t *testing.T) {
	i := []int{1, 9, 10, 3, 99, 3, 11, 0, 0, 30, 40, 50}
	cpu := CPU{intcode: i}

	// Get an ADD -- ignore
	b := cpu.DoNext()
	if !b {
		t.Errorf("DoNext returned false")
	}

	// Get a HALT -- return true but set halted flag
	b = cpu.DoNext()
	if !b {
		t.Errorf("DoNext returned false")
	}

	// Should now be halted (and return false)
	b = cpu.DoNext()
	if b {
		t.Errorf("DoNext returned true")
	}

	if cpu.Value() != 0 {
		t.Errorf("bad value: expected=0, got=%d\n", cpu.Value())
	}
}

func TestMul(t *testing.T) {
	i := []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	cpu := CPU{intcode: i, ip: 4}

	t.Log("IN: ", cpu.intcode)

	peek := cpu.Peek()
	if !SlicesEqual(peek, i[4:8]) {
		t.Fatal("Unexpected read order:", peek)
	}

	b := cpu.DoNext()

	if !b {
		t.Errorf("DoNext returned false")
	}
	if i[0] != 3500 {
		t.Errorf("add failed: expected 3500, got %d\n", i[0])
	}

	if cpu.Value() != 3500 {
		t.Errorf("bad value: expected=3500, got=%d\n", cpu.Value())
	}

	t.Log("OUT:", cpu.intcode)
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
