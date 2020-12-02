package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	testLine := "1-13 a: abcde"

	l, err := parseLine(testLine)
	if err != nil {
		t.Error(err)
	}

	if l.min != 1 || l.max != 13 || l.letter != "a" || l.password != "abcde" {
		t.Error(l)
	}
}
