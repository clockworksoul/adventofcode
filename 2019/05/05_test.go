package main

import (
	"fmt"
	"testing"
)

func TestReadNormal(t *testing.T) {
	intcode := []int{1002, 4, 3, 4, 33}

	cpu := NewCPU(intcode)
	t.Log("IN: ", cpu.intcode)

	opcode, modes, params, err := cpu.Peek()
	if err != nil {
		t.Fatal(err)
	}

	if opcode != CodeMul {
		t.Errorf("bad opcode: expected %d, got %d\n", CodeMul, opcode)
	}

	if len(params) != 3 {
		t.Fatal("unexpected number of params:", len(params))
	}
	if len(modes) != 3 {
		t.Fatal("unexpected number of modes:", len(modes))
	}

	if !SlicesEqual(params, []int{4, 3, 4}) {
		t.Fatal("unexpected read order:", params)
	}
	if modes[0] != ModePosition ||
		modes[1] != ModeImmediate ||
		modes[2] != ModePosition {

		t.Fatal("unexpected mode: got=", modes)
	}

	t.Log("A: ", cpu)
	err = cpu.Execute()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("B: ", cpu)

	if cpu.intcode[4] != 99 {
		t.Errorf("bad output: expected %d, got %d\n", 99, cpu.intcode[4])
	}
}

func TestInOut(t *testing.T) {
	const input = 1234
	intcode := []int{3, 0, 4, 0, 99}

	cpu := NewCPU(intcode)
	fmt.Println("START: ", cpu.intcode)

	cpu.istream <- input

	fmt.Println("IN   : ", cpu.intcode)

	output := 0

	go func() {
		for {
			select {
			case output = <-cpu.ostream:
				t.Log(output)
			case <-cpu.halt:
			}
		}
	}()

	err := cpu.ExecuteAll()
	if err != nil {
		t.Error(err)
	}

	if input != output {
		t.Errorf("bad output: expected %d, got %d\n", input, output)
	}
}

func TestJumpPosMode(t *testing.T) {
	intcode := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	cpu := NewCPU(intcode)
	output := 0

	cpu.istream <- 8

	go func() {
		for {
			select {
			case output = <-cpu.ostream:
				t.Log(output)
			case <-cpu.halt:
				t.Log("Done")
			}
		}
	}()

	err := cpu.ExecuteAll()
	if err != nil {
		t.Error(err)
	}
}

// func TestPeekOverrun(t *testing.T) {
// 	i := []int{1, 9, 10, 3}
// 	cpu := CPU{intcode: i, ip: 4}

// 	peek := cpu.PeekN(4)
// 	if !SlicesEqual(peek, []int{}) {
// 		t.Fatal("unexpected read order:", peek)
// 	}
// }

// func TestAdd(t *testing.T) {
// 	i := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
// 	cpu := CPU{intcode: i}

// 	t.Log("IN: ", cpu.intcode)

// 	peek := cpu.Peek()
// 	if !SlicesEqual(peek, i[0:4]) {
// 		t.Fatal("unexpected read order:", peek)
// 	}

// 	b := cpu.Execute()

// 	if !b {
// 		t.Errorf("DoNext returned false")
// 	}
// 	if i[3] != 70 {
// 		t.Errorf("add failed: expected 70, got %d\n", i[3])
// 	}

// 	t.Log("OUT:", cpu.intcode)
// }

// func TestHalt(t *testing.T) {
// 	i := []int{1, 9, 10, 3, 99, 3, 11, 0, 0, 30, 40, 50}
// 	cpu := CPU{intcode: i}

// 	// Get an ADD -- ignore
// 	b := cpu.DoNext()
// 	if !b {
// 		t.Errorf("DoNext returned false")
// 	}

// 	// Get a HALT -- return true but set halted flag
// 	b = cpu.DoNext()
// 	if !b {
// 		t.Errorf("DoNext returned false")
// 	}

// 	// Should now be halted (and return false)
// 	b = cpu.DoNext()
// 	if b {
// 		t.Errorf("DoNext returned true")
// 	}
// }

// func TestMultiply(t *testing.T) {
// 	intcode := []int{1002, 4, 3, 4, 33}
// 	cpu := NewCPU(intcode)

// 	t.Log("IN: ", cpu.intcode)

// 	err := cpu.Execute()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if cpu.Value() != 3500 {
// 		t.Errorf("add failed: expected 3500, got %d\n", i[0])
// 	}

// 	t.Log("OUT:", cpu.intcode)
// }

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
