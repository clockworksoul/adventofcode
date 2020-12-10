package main

import (
	"fmt"
	"log"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	inputs := []int{}

	if err := adventofcode.IngestFileInt("input.txt", func(i int) {
		inputs = append(inputs, i)
	}); err != nil {
		log.Fatal(err)
	}

	for _, i := range inputs {
		for _, j := range inputs {
			for _, k := range inputs {
				if i+j+k == 2020 {
					fmt.Printf("%d + %d + %d = 2020 -> %d\n", i, j, k, i*j*k)
				}
			}
		}
	}
}
