package main

import (
	"fmt"
)

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
