package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

type Seat rune

const (
	Floor    Seat = ','
	Empty    Seat = 'L'
	Occupied Seat = '#'
)

func main() {
	var seats = [][]Seat{}

	adventofcode.IngestFile("input.txt", func(line string) {
		seats = append(seats, IngestLine(line))
	})

	StarOne(seats)
	StarTwo(seats)
}

func StarOne(seats [][]Seat) {
	seats, changes, occupied := Tick(seats, 4, CountOccupiedOne)
	for changes > 0 {
		seats, changes, occupied = Tick(seats, 4, CountOccupiedOne)
	}
	fmt.Println("One:", occupied)
}

func StarTwo(seats [][]Seat) {
	seats, changes, occupied := Tick(seats, 5, CountOccupiedTwo)
	for changes > 0 {
		seats, changes, occupied = Tick(seats, 5, CountOccupiedTwo)
	}
	fmt.Println("Two:", occupied)
}

func Tick(seats [][]Seat, occupiedThreshold int, count func(x, y int, s [][]Seat) int) ([][]Seat, int, int) {
	changes, occupied := 0, 0
	ng := [][]Seat{}
	for _, s := range seats {
		ng = append(ng, make([]Seat, len(s)))
	}

	for y := range seats {
		for x := range seats[y] {
			sum := count(x, y, seats)

			switch {
			case seats[y][x] == Empty && sum == 0:
				ng[y][x] = Occupied
			case seats[y][x] == Occupied && sum >= occupiedThreshold:
				ng[y][x] = Empty
			default:
				ng[y][x] = seats[y][x]
			}

			if seats[y][x] != ng[y][x] {
				changes++
			}
			if ng[y][x] == Occupied {
				occupied++
			}
		}
	}

	return ng, changes, occupied
}

func CountOccupiedOne(x, y int, s [][]Seat) int {
	count := func(x, y int, s [][]Seat) int {
		if SeatValid(x, y, s) && s[y][x] == Occupied {
			return 1
		}
		return 0
	}

	sum := count(x-1, y-1, s) +
		count(x-0, y-1, s) +
		count(x+1, y-1, s) +
		count(x-1, y-0, s) +
		count(x+1, y-0, s) +
		count(x-1, y+1, s) +
		count(x-0, y+1, s) +
		count(x+1, y+1, s)

	return sum
}

func CountOccupiedTwo(x, y int, s [][]Seat) int {
	sum := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			for sx, sy := x+dx, y+dy; SeatValid(sx, sy, s); sx, sy = sx+dx, sy+dy {
				if s[sy][sx] == Occupied {
					sum++
					break
				} else if s[sy][sx] == Empty {
					break
				}
			}
		}
	}

	return sum
}

func SeatValid(x, y int, s [][]Seat) bool {
	return y >= 0 && y < len(s) && x >= 0 && x < len(s[y])
}

func IngestLine(line string) []Seat {
	s := make([]Seat, len(line))
	for i, c := range line {
		s[i] = Seat(c)
	}
	return s
}
