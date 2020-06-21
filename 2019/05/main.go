package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	intcode, err := ingestInputs()
	cpu := NewCPU(intcode)

	cpu.istream <- 5

	go func() {
		for {
			select {
			case out := <-cpu.ostream:
				fmt.Println(out)
			case <-cpu.halt:
				fmt.Println("Done")
			}
		}
	}()

	err = cpu.ExecuteAll()
	if err != nil {
		panic(err)
	}
}

func ingestInputs() ([]int, error) {
	ints := make([]int, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("cannot open input file: %w", err)
	}

	str := ""
	reader := bufio.NewReader(file)

	for err == nil {
		str, err = reader.ReadString(',')
		i := 0

		fmt.Sscanf(strings.Trim(str, ",\n"), "%d", &i)
		ints = append(ints, i)
	}

	return ints, nil
}
