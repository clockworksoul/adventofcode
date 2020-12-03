package main

import (
	"adventofcode"
	"fmt"
	"log"
	"strconv"
)

func main() {
	inputs, err := scanInput()
	if err != nil {
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

func scanInput() ([]int, error) {
	var inputs = []int{}

	adventofcode.IngestFileE("./input.txt", func(txt string) error {
		i, err := strconv.Atoi(txt)
		if err != nil {
			return fmt.Errorf("warning: cannot parse %s as integer", txt)
		}

		inputs = append(inputs, i)
		return nil
	})

	return inputs, nil
}
