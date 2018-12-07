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

// Assumes unit has a len of 1
func ScrubUnits(polymer string, unit string) string {
	runeL := rune(strings.ToLower(unit)[0])
	runeU := rune(strings.ToUpper(unit)[0])

	newpolymer := make([]rune, 0)

	for _, r := range polymer {
		if r != runeL && r != runeU {
			newpolymer = append(newpolymer, r)
		}
	}

	return string(newpolymer)
}

func FindShortestWithOneUnitRemoved(polymer string) (int, string) {
	runes := make(map[string]int)

	leastRune := "🔥"
	leastLen := len(polymer)

	for i, _ := range polymer {
		runeU := strings.ToUpper(polymer[i : i+1])

		if _, ok := runes[runeU]; !ok {
			scrubbed := ScrubUnits(polymer, runeU)
			collapsed, _ := CollapsePolymer(scrubbed)
			runes[runeU] = len(collapsed)

			if len(collapsed) < leastLen {
				leastRune = runeU
				leastLen = len(collapsed)
			}
		}
	}

	return leastLen, leastRune
}

func main() {
	polymer, err := IngestPolymer("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	collapsed, _ := CollapsePolymer(polymer)
	fmt.Println(len(collapsed))

	length, unit := FindShortestWithOneUnitRemoved(polymer)
	fmt.Println(length, unit)
}
