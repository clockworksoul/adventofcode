package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var trees []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
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
