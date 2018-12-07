package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IngestPolymer(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", nil
}

// Collapses the polymer.
// Returns the new polymer and the number of rounds required to collapse
func CollapsePolymer(polymer string) (string, int) {
	count := 1
	totalModifications := 0

	for count > 0 {
		count = 0

		for i := 0; i < len(polymer)-2; {
			a := polymer[i : i+1]
			b := polymer[i+1 : i+2]

			if a != b && strings.ToUpper(a) == strings.ToUpper(b) {
				polymer = polymer[:i] + polymer[i+2:]
				count++
			} else {
				i++
			}
		}

		totalModifications += count
	}

	return polymer, totalModifications
}

func main() {
	polymer, err := IngestPolymer("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	collapsed, _ := CollapsePolymer(polymer)

	fmt.Println(len(collapsed))
}
