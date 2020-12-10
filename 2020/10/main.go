package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/clockworksoul/adventofcode"
)

var cache = make(map[int]int)

func main() {
	adapters := make([]int, 0)

	if err := adventofcode.IngestFileInt("input.txt", func(i int) {
		adapters = append(adapters, i)
	}); err != nil {
		log.Fatal(err)
	} else {
		sort.Ints(adapters)
		adapters = append(adapters, 3+adapters[len(adapters)-1])

		fmt.Println("Part 1:", partOne(adapters))
		fmt.Println("Part 2:", partTwo(0, adapters[len(adapters)-1], adapters))
	}
}

func partOne(adapters []int) int {
	j1, j3, jolts := 0, 0, 0
	for _, j := range adapters {
		switch j - jolts {
		case 1:
			j1++
		case 3:
			j3++
		default:
			log.Println("Oops! Got ", j-jolts)
		}
		jolts = j
	}

	return j1 * j3
}

func partTwo(cv, goal int, adapters []int) int {
	if result, ok := cache[cv]; ok {
		return result
	} else if cv == goal {
		cache[cv] = 1
	} else {
		for i, v := range adapters {
			if v > cv && v <= cv+3 {
				result += partTwo(v, goal, adapters[i:])
			} else if v > cv+3 {
				break
			}
		}
		cache[cv] = result
	}

	return cache[cv]
}
