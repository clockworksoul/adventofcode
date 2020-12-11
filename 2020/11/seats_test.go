package main

import (
	"fmt"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestSeats8(t *testing.T) {
	text := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`
	seats := textToSeats(text)
	sum := CountOccupiedTwo(3, 4, seats)

	assert.Equal(t, 8, sum)
}

func TestSeats1(t *testing.T) {
	text := `.............
.L.L.#.#.#.#.
.............`
	seats := textToSeats(text)
	sum := CountOccupiedTwo(1, 1, seats)

	assert.Equal(t, 0, sum)
}

func TestSeats0(t *testing.T) {
	text := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	seats := textToSeats(text)
	sum := CountOccupiedTwo(3, 3, seats)

	assert.Equal(t, 0, sum)
}

func textToSeats(text string) [][]Seat {
	seats := [][]Seat{}
	lines := strings.Split(text, "\n")

	for _, s := range lines {
		seats = append(seats, IngestLine(s))
	}

	return seats
}

func Print(seats [][]Seat) {
	line := ""
	for _, y := range seats {
		line = ""
		for _, s := range y {
			line += string(s)
		}
		fmt.Println(line)
	}
}
