package main

import "fmt"

var (
	input = []int{9, 3, 1, 0, 8, 4}
)

func main() {
	var spoken []int

	spoken = speak(input, 2020)
	fmt.Println("Star One:", spoken[len(spoken)-1])

	spoken = speak(input, 30000000)
	fmt.Println("Star Two:", spoken[len(spoken)-1])
}

func speak(input []int, max int) []int {
	speak := []int{}
	spoken := map[int][]int{}
	last := 0

	for i, v := range input {
		last = v
		speak = append(speak, v)
		spoken[v] = append(spoken[v], i)
	}

	for i := len(input); i < max; i++ {
		before, ok := spoken[last]

		if !ok || len(before) == 1 {
			last = 0
			speak = append(speak, last)
			spoken[last] = append(spoken[last], i)
		} else {
			l := len(before)
			last = before[l-1] - before[l-2]
			speak = append(speak, last)
			spoken[last] = append(spoken[last], i)
		}
	}

	return speak
}
