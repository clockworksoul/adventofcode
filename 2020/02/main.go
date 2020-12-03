package main

import (
	"adventofcode"
	"fmt"
	"log"
	"strings"
)

func main() {
	count := 0

	err := adventofcode.IngestFileE("./input.txt", func(txt string) error {
		l, err := parseLine(txt)
		if err != nil {
			return fmt.Errorf("parse failure: \"%s\"", txt)
		}

		if evaluateLine2(l) {
			count++
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

type LineParts struct {
	min      int
	max      int
	letter   string
	password string
}

func parseLine(line string) (LineParts, error) {
	l := LineParts{}
	_, err := fmt.Sscanf(line, "%d-%d %s %s", &l.min, &l.max, &l.letter, &l.password)
	l.letter = string(l.letter[0])
	return l, err
}

func evaluateLine1(l LineParts) bool {
	count := strings.Count(l.password, l.letter)

	return count >= l.min && count <= l.max
}

func evaluateLine2(l LineParts) bool {
	l1 := string(l.password[l.min-1]) == l.letter
	l2 := string(l.password[l.max-1]) == l.letter

	return (l1 || l2) && !(l1 && l2)
}
