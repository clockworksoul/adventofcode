package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	inputs := []int{}

	adventofcode.IngestFileInt("input.txt", func(i int) {
		inputs = append(inputs, i)
	})

	fmt.Println("WINDOW 1:", countIncreases(inputs, 1))
	fmt.Println("WINDOW 3:", countIncreases(inputs, 3))
}

func countIncreases(inputs []int, window int) int {
	lastInput := sumWindow(inputs, 0, window)
	increases := 0

	for i := window + 1; i <= len(inputs); i++ {
		input := sumWindow(inputs, i-window, window)
		if input > lastInput {
			increases++
		}

		lastInput = input
	}

	return increases
}

func sumWindow(inputs []int, start, window int) int {
	sum := 0

	for i := 0; i < window; i++ {
		sum += inputs[i+start]
	}

	return sum
}
