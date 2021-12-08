package main

import (
	"github.com/clockworksoul/adventofcode"
)

func main() {
	var inputs []string
	adventofcode.IngestFile("example.txt", func(s string) {
		inputs = append(inputs, s)
	})

	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
}

func part2(inputs []string) {
}
