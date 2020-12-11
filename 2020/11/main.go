package main

import (
	"fmt"

	"github.com/clockworksoul/adventofcode"
)

type seat rune

const (
	floor    seat = ','
	empty    seat = 'L'
	occupied seat = '#'
)

func main() {
	var seats = [][]seat{}

	adventofcode.IngestFile("input.txt", func(line string) {
		seats = append(seats, ingestLine(line))
	})

	starOne(seats)
	starTwo(seats)
}

func starOne(seats [][]seat) {
	seats, changes, occupied := tick(seats, 4, countOccupiedOne)
	for changes > 0 {
		seats, changes, occupied = tick(seats, 4, countOccupiedOne)
	}
	fmt.Println("One:", occupied)
}

func starTwo(seats [][]seat) {
	seats, changes, occupied := tick(seats, 5, countOccupiedTwo)
	for changes > 0 {
		seats, changes, occupied = tick(seats, 5, countOccupiedTwo)
	}
	fmt.Println("Two:", occupied)
}

func tick(seats [][]seat, occupiedThreshold int, count func(x, y int, s [][]seat) int) ([][]seat, int, int) {
	changes, occ := 0, 0
	ng := [][]seat{}
	for _, s := range seats {
		ng = append(ng, make([]seat, len(s)))
	}

	for y := range seats {
		for x := range seats[y] {
			sum := count(x, y, seats)

			switch {
			case seats[y][x] == empty && sum == 0:
				ng[y][x] = occupied
			case seats[y][x] == occupied && sum >= occupiedThreshold:
				ng[y][x] = empty
			default:
				ng[y][x] = seats[y][x]
			}

			if seats[y][x] != ng[y][x] {
				changes++
			}
			if ng[y][x] == occupied {
				occ++
			}
		}
	}

	return ng, changes, occ
}

func countOccupiedOne(x, y int, s [][]seat) int {
	count := func(x, y int, s [][]seat) int {
		if seatValid(x, y, s) && s[y][x] == occupied {
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

func countOccupiedTwo(x, y int, s [][]seat) int {
	sum := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			for sx, sy := x+dx, y+dy; seatValid(sx, sy, s); sx, sy = sx+dx, sy+dy {
				if s[sy][sx] == occupied {
					sum++
					break
				} else if s[sy][sx] == empty {
					break
				}
			}
		}
	}

	return sum
}

func seatValid(x, y int, s [][]seat) bool {
	return y >= 0 && y < len(s) && x >= 0 && x < len(s[y])
}

func ingestLine(line string) []seat {
	s := make([]seat, len(line))
	for i, c := range line {
		s[i] = seat(c)
	}
	return s
}
