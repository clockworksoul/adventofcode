package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	CodeAdd  int = 1
	CodeMul  int = 2
	CodeHalt int = 99
)

func main() {
	ints, err := ingestInputs()
	if err != nil {
		log.Fatal(err)
	}

	// before running the program, replace position 1 with the value 12 and
	// replace position 2 with the value 2. 
	ints[1] = 12
	ints[2] = 2

	cpu := CPU{intcode:ints}

	for cpu.DoNext() {
		log.Println("Processing...")
	}

	// What value is left at position 0 after the program halts?
	log.Println("Value at position 0:", cpu.intcode[0])
}

func ingestInputs() ([]int, error) {
	ints := make([]int, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("cannot open input file: %w", err)
	}

	str := ""
	reader := bufio.NewReader(file)

	for err == nil {
		str, err = reader.ReadString(',')
		i := 0

		fmt.Sscanf(strings.Trim(str, ",\n"), "%d", &i)
		ints = append(ints, i)
	}

	return ints, nil
}

type CPU struct {
	intcode []int
	ip      int
	halted  bool
}

func (cpu *CPU) DoNext() bool {
	if cpu.halted || cpu.ip >= len(cpu.intcode) {
		return false
	}

	switch cpu.intcode[cpu.ip] {
	case CodeAdd:
		cpu.executeAdd()
	case CodeMul:
		cpu.executeMul()
	case CodeHalt:
		cpu.halted = true
	}

	cpu.ip += 4

	return true
}

func (cpu *CPU) Peek() []int {
	ints := make([]int, 4)

	for i := 0; i < 4 && i + cpu.ip < len(cpu.intcode); i++ {
		ints[i] = cpu.intcode[cpu.ip + i]
	}

	return ints
}

func (cpu *CPU) executeAdd() {
	peek := cpu.Peek()
	pa, pb, pc := peek[1], peek[2], peek[3]
	a, b := cpu.intcode[pa], cpu.intcode[pb]
	cpu.intcode[pc] = a + b
}

func (cpu *CPU) executeMul() {
	peek := cpu.Peek()
	pa, pb, pc := peek[1], peek[2], peek[3]
	a, b := cpu.intcode[pa], cpu.intcode[pb]
	cpu.intcode[pc] = a * b
}
