package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	part1()
	part2()
}

func part1() int {
	var x, depth int

	adventofcode.IngestFile("input.txt", func(s string) {
		var dir string
		var amt int

		fmt.Sscanf(s, "%s %d", &dir, &amt)

		switch dir {
		case "forward":
			x += amt
		case "up":
			depth -= amt
		case "down":
			depth += amt
		default:
			panic(dir)
		}
	})

	return x * depth
}

func part2() int {
	var x, depth, aim int

	adventofcode.IngestFile("input.txt", func(s string) {
		var dir string
		var amt int

		fmt.Sscanf(s, "%s %d", &dir, &amt)

		switch dir {
		case "forward":
			x += amt
			depth += aim * amt
		case "up":
			aim -= amt
		case "down":
			aim += amt
		default:
			panic(dir)
		}
	})

	return x * depth
}
