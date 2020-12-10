package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/clockworksoul/adventofcode"
)

var cache = make(map[int]int)

func main() {
	adapters := []int{}

	adventofcode.IngestFile("./input.txt", func(line string) {
		i, _ := strconv.Atoi(line)
		adapters = append(adapters, i)
	})

	sort.Ints(adapters)
	adapters = append(adapters, 3+adapters[len(adapters)-1])

	partOne(adapters)
	partTwo(adapters)
}

func partOne(adapters []int) {
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

	fmt.Println("Part 1:", j1*j3)
}

func partTwo(adapters []int) {
	fmt.Println("Part 2:", seek(0, adapters[len(adapters)-1], adapters))
}

func seek(cv, goal int, adapters []int) int {
	if result, ok := cache[cv]; ok {
		return result
	} else if cv == goal {
		cache[cv] = 1
	} else {
		for i, v := range adapters {
			if v > cv && v <= cv+3 {
				result += seek(v, goal, adapters[i:])
			} else if v > cv+3 {
				break
			}
		}
		cache[cv] = result
	}

	return cache[cv]
}
