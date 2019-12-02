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
	log.Println("(1) Value at position 0:", cpu.Value())
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

			if cpu.Value() == seeking {
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
}

func (cpu *CPU) DoNext() bool {
	if cpu.halted || cpu.ip >= len(cpu.intcode) {
		return false
	}

	var instructionLength int

	switch cpu.Peek() {
	case CodeAdd:
		instructionLength = cpu.executeAdd()
	case CodeMul:
		instructionLength = cpu.executeMul()
	case CodeHalt:
		instructionLength = cpu.executeHalt()
	}

	cpu.ip += instructionLength

	return true
}

func (cpu *CPU) Peek() int {
	if cpu.ip >= len(cpu.intcode) {
		return 0
	}

	return cpu.intcode[cpu.ip]	
}

func (cpu *CPU) PeekN(length int) []int {
	start, end := cpu.ip, cpu.ip + length

	if end > len(cpu.intcode) {
		end = len(cpu.intcode)
	}

	return cpu.intcode[start:end]
}

func (cpu *CPU) Value() int {
	return cpu.intcode[0]
}

func (cpu *CPU) executeAdd() int {
	peek := cpu.PeekN(4)
	pnoun, pverb, paddress := peek[1], peek[2], peek[3]
	noun, verb := cpu.intcode[pnoun], cpu.intcode[pverb]
	cpu.intcode[paddress] = noun + verb
	return 4
}

func (cpu *CPU) executeMul() int {
	peek := cpu.PeekN(4)
	pnoun, pverb, paddress := peek[1], peek[2], peek[3]
	noun, verb := cpu.intcode[pnoun], cpu.intcode[pverb]
	cpu.intcode[paddress] = noun * verb
	return 4
}

func (cpu *CPU) executeHalt() int {
	cpu.halted = true
	return 1
}
