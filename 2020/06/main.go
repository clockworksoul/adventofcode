package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	current := map[string]int{}
	sum, count := 0, 0

	adventofcode.IngestFile("./input.txt", func(txt string) {
		if txt == "" {
			for _, v := range current {
				if v == count {
					sum++
				}
			}

			current = map[string]int{}
			count = 0
		} else {
			count++
			for _, ch := range txt {
				current[string(ch)]++
			}
		}
	})

	for _, v := range current {
		if v == count {
			sum++
		}
	}

	fmt.Println(sum)
}
