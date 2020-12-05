package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	present := make([]bool, 1<<10)
	highest := 0

	adventofcode.IngestFile("./input.txt", func(txt string) {
		row, column := 0, 0

		for i, f, b := 0, 0, 127; i < 8; i++ {
			switch txt[i] {
			case 'F':
				next := (f + b) / 2
				b, row = next, next
			case 'B':
				next := (f + b + 1) / 2
				f, row = next, next
			}
		}

		for i, l, r := 7, 0, 7; i < 10; i++ {
			switch txt[i] {
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
			fmt.Println("Missing: ", i)
		}
	}
}
