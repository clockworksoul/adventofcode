package main

import (
	"adventofcode"
	"fmt"
	"log"
)

func main() {
	trees := []string{}

	err := adventofcode.IngestFile("./input.txt", func(txt string) {
		trees = append(trees, txt)
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countTrees(trees, 1, 1) *
		countTrees(trees, 3, 1) *
		countTrees(trees, 5, 1) *
		countTrees(trees, 7, 1) *
		countTrees(trees, 1, 2))
}

func countTrees(trees []string, dx, dy int) int {
	count := 0

	for x, y := 0, 0; y < len(trees); x, y = x+dx, y+dy {
		if trees[y][x%len(trees[y])] == '#' {
			count++
		}
	}

	return count
}
