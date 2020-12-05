package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	present := make([]bool, 1<<10)
	highest := 0

	adventofcode.IngestFile("./input.txt", func(txt string) {
		var row, column int

		for i, f, b, l, r := 0, 0, 127, 0, 7; i < 10; i++ {
			switch txt[i] {
			case 'F':
				next := (f + b) / 2
				b, row = next, next
			case 'B':
				next := (f + b + 1) / 2
				f, row = next, next
			case 'L':
				next := (l + r) / 2
				r, column = next, next
			case 'R':
				next := (l + r + 1) / 2
				l, column = next, next
			}
		}

		seat := (row * 8) + column
		if seat > highest {
			highest = seat
		}

		present[seat] = true
	})

	fmt.Println("Highest:", highest)

	for i := 100; i < len(present)-100; i++ {
		if !present[i] {
			fmt.Println("Missing:", i)
		}
	}
}
