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

	One(seats)
	Two(seats)
}

func IngestLine(line string) []Seat {
	s := make([]Seat, len(line))
	for i, c := range line {
		s[i] = Seat(c)
	}
	return s
}

func One(seats [][]Seat) {
	seats, changes, occupied := Tick1(seats)
	for changes > 0 {
		seats, changes, occupied = Tick1(seats)
	}
	fmt.Println("One:", occupied)
}

func Two(seats [][]Seat) {
	seats, changes, occupied := Tick2(seats)
	for changes > 0 {
		seats, changes, occupied = Tick2(seats)
	}
	fmt.Println("Two:", occupied)
}

func Tick1(seats [][]Seat) ([][]Seat, int, int) {
	changes, occupied := 0, 0
	ng := [][]Seat{}
	for _, s := range seats {
		ng = append(ng, make([]Seat, len(s)))
	}

	for y := range seats {
		for x := range seats[y] {
			sum := count(x-1, y-1, seats) +
				count(x-0, y-1, seats) +
				count(x+1, y-1, seats) +
				count(x-1, y-0, seats) +
				count(x+1, y-0, seats) +
				count(x-1, y+1, seats) +
				count(x-0, y+1, seats) +
				count(x+1, y+1, seats)

			switch {
			case seats[y][x] == Empty && sum == 0:
				ng[y][x] = Occupied
			case seats[y][x] == Occupied && sum >= 4:
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

func Tick2(seats [][]Seat) ([][]Seat, int, int) {
	changes, occupied := 0, 0
	ng := [][]Seat{}
	for _, s := range seats {
		ng = append(ng, make([]Seat, len(s)))
	}

	for y := range seats {
		for x := range seats[y] {
			sum := CountOccupied2(x, y, seats)

			switch {
			case seats[y][x] == Empty && sum == 0:
				ng[y][x] = Occupied
			case seats[y][x] == Occupied && sum >= 5:
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

func CountOccupied2(x, y int, s [][]Seat) int {
	sum := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			for sx, sy := x+dx, y+dy; valid(sx, sy, s); sx, sy = sx+dx, sy+dy {
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

func valid(x, y int, s [][]Seat) bool {
	return y >= 0 && y < len(s) && x >= 0 && x < len(s[y])
}

func count(x, y int, s [][]Seat) int {
	if valid(x, y, s) && s[y][x] == Occupied {
		return 1
	}
	return 0
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
