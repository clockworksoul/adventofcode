package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	numbers := []int{}

	adventofcode.IngestFile("./input.txt", func(line string) {
		i, _ := strconv.Atoi(line)
		numbers = append(numbers, i)
	})

	invalid := first(numbers, 25)
	fmt.Println(invalid)

	cont := sumTo(numbers, invalid)
	fmt.Println(cont)

	sort.Ints(cont)
	fmt.Println(cont[0] + cont[len(cont)-1])
}

func first(numbers []int, length int) int {

outer:
	for i := length; i < len(numbers); i++ {
		preamble := numbers[i-length : i]
		for _, x := range preamble {
			for _, y := range preamble {
				if x+y == numbers[i] && x != y {
					continue outer
				}
			}
		}
		return numbers[i]
	}

	return 0
}

func sumTo(numbers []int, n int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] > n {
				break
			} else if n == sum(numbers[i:j]...) {
				return numbers[i:j]
			}
		}
	}

	return []int{}
}

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
