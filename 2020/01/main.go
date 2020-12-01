package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	file, err := os.Open("./input.txt")
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		i, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Printf("warning: cannot parse %s as integer", txt)
		}

		inputs = append(inputs, i)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	return inputs, nil
}
