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
	day1()
	day2()
}

func day1() {
	ints, err := ingestInputs()
	if err != nil {
		log.Fatal(err)
	}

	// Before running the program, replace position 1 with the value 12 and
	// replace position 2 with the value 2.
	ints[1] = 12
	ints[2] = 2

	cpu := CPU{intcode: ints}

	for cpu.DoNext() {
		// Processing...
	}

	// What value is left at position 0 after the program halts?
	log.Println("(1) Value at position 0:", cpu.intcode[0])
}

func day2() {
	const seeking = 19690720

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ints, err := ingestInputs()
			if err != nil {
				log.Fatal(err)
			}

			ints[1] = noun
			ints[2] = verb

			cpu := CPU{intcode: ints}

			for cpu.DoNext() {
				// Processing...
			}

			if cpu.intcode[0] == seeking {
				log.Printf("(2) Found %d! noun=%d; verb=%d\n", seeking, noun, verb)
				log.Printf("(2) 100 * noun + verb = %d\n", (100*noun)+verb)
				return
			}
		}
	}

	log.Println("(2) No match found.")
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
	value   int
}

func (cpu *CPU) DoNext() bool {
	var val int

	if cpu.halted || cpu.ip >= len(cpu.intcode) {
		return false
	}

	switch cpu.intcode[cpu.ip] {
	case CodeAdd:
		val = cpu.executeAdd()
		cpu.ip += 4
	case CodeMul:
		val = cpu.executeMul()
		cpu.ip += 4
	case CodeHalt:
		cpu.halted = true
		cpu.ip += 1
	}

	cpu.value = val

	return true
}

func (cpu *CPU) Value() int {
	return cpu.value
}

func (cpu *CPU) Peek() []int {
	ints := make([]int, 4)

	for i := 0; i < 4 && i+cpu.ip < len(cpu.intcode); i++ {
		ints[i] = cpu.intcode[cpu.ip+i]
	}

	return ints
}

func (cpu *CPU) executeAdd() int {
	peek := cpu.Peek()
	pnoun, pverb, paddress := peek[1], peek[2], peek[3]
	noun, verb := cpu.intcode[pnoun], cpu.intcode[pverb]
	cpu.intcode[paddress] = noun + verb

	return cpu.intcode[paddress]
}

func (cpu *CPU) executeMul() int {
	peek := cpu.Peek()
	pnoun, pverb, paddress := peek[1], peek[2], peek[3]
	noun, verb := cpu.intcode[pnoun], cpu.intcode[pverb]
	cpu.intcode[paddress] = noun * verb

	return cpu.intcode[paddress]
}
