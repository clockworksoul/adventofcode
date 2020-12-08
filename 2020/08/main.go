package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

type Op struct {
	Code  string
	Value int
}

var instructions = make([]*Op, 0)

func main() {
	adventofcode.IngestFile("./input.txt", func(line string) {
		op := Op{}
		fmt.Sscanf(line, "%s %d", &op.Code, &op.Value)
		instructions = append(instructions, &op)
	})

	res := 0
	for i := 0; res != resultTerminated && i < len(instructions); i++ {
		res = execute(i)
		switch res {
		case resultInfiniteLoop:
			fmt.Println(i, "Infinite loop!")
		case resultTerminated:
			fmt.Println(i, "Terminated normally!")
		}
	}
}

const (
	resultInfiniteLoop int = iota
	resultTerminated
)

func execute(mutate int) int {
	var acc, ip int
	code, value := "", 0

	counts := make([]int, len(instructions))

	for ip < len(instructions) {
		code, value = instructions[ip].Code, instructions[ip].Value
		counts[ip]++
		count := counts[ip]

		if ip == mutate {
			if code == "jmp" {
				code = "nop"
			} else if code == "nop" {
				code = "jmp"
			}
		}

		switch code {
		case "acc":
			acc += value
			ip++
		case "jmp":
			ip += value
		case "nop":
			ip++
		default:
			panic("Unknown code: " + code)
		}

		if count > 2 {
			return resultInfiniteLoop
		}
	}

	fmt.Printf("%s %+d --> %d\n", code, value, acc)

	return resultTerminated
}
