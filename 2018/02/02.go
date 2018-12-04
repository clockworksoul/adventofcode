package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	biggestCount = 3
)

func readStrings(filename string) (chan string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	c := make(chan string)
	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			c <- scanner.Text()
		}

		close(c)
	}()

	return c, nil
}

func hasNMultiples(id string, max int) []bool {
	countSlice := make([]bool, max+1)
	countMap := make(map[rune]int)

	for _, r := range id {
		countMap[r]++
	}

	for _, v := range countMap {
		if v < len(countSlice) {
			countSlice[v] = true
		}
	}

	return countSlice
}

func calculateChecksum(filename string) int {
	ch, err := readStrings(filename)
	if err != nil {
		log.Fatal(err)
	}

	countsAccumulator := make([]int, biggestCount+1)

	for str := range ch {
		multiples := hasNMultiples(str, biggestCount)
		for i := 2; i < len(multiples); i++ {
			if multiples[i] {
				countsAccumulator[i]++
			}
		}
	}

	sum := 0

	for _, count := range countsAccumulator {
		if count > 0 {
			if sum == 0 {
				sum = count
			} else {
				sum *= count
			}
		}
	}

	return sum
}

// Assume that strings have the same length
// Return letters in common, count of differences
func diffStrings(a, b string) (string, int) {
	count := 0
	same := ""

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			same += string(a[i])
		} else {
			count++
		}
	}

	return same, count
}

func findDifferences(filename string) string {
	ch, err := readStrings(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Lazy but yeah. Read in all the strings.
	ids := make([]string, 0)
	for id := range ch {
		ids = append(ids, id)
	}

	for i, id1 := range ids {
		for j := i + 1; j < len(ids); j++ {
			id2 := ids[j]
			same, diff := diffStrings(id1, id2)

			if diff == 1 {
				return same
			}
		}
	}

	return ""
}

func main() {
	solution1 := calculateChecksum("input.txt")
	fmt.Println(solution1)

	solution2 := findDifferences("input.txt")
	fmt.Println(solution2)
}
