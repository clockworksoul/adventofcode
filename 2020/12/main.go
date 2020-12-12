package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/clockworksoul/adventofcode"
)

var (
	directions = [4][2]int{
		{1, 0},  // east
		{0, -1}, // south
		{-1, 0}, // west
		{0, 1}}  // north
)

const (
	left  int = -1
	right int = 1

	ew = 0
	ns = 1

	east  = 0
	south = 1
	west  = 2
	north = 3
)

type position struct {
	pos    [2]int
	facing int
}

func (p *position) move(direction int, units int) {
	p.pos[0] += directions[direction][0] * units
	p.pos[1] += directions[direction][1] * units
}

func (p *position) turn(facing int, degrees int) {
	p.facing += (facing * (degrees / 90))
	p.facing %= 4
	if p.facing < 0 {
		p.facing = (4 + p.facing) % 4
	}
}

func (p *position) forward(units int) {
	p.pos[0] += directions[p.facing][0] * units
	p.pos[1] += directions[p.facing][1] * units
}

func (p *position) manhattan() int {
	m := math.Abs(float64(p.pos[0])) + math.Abs(float64(p.pos[1]))
	return int(m)
}

func main() {
	p := &position{waypoint: [2]int{0, 0}}

	adventofcode.IngestFile("input.txt", func(line string) {
		fmt.Print(line, " ")
		c := line[0]
		units, _ := strconv.Atoi(line[1:])

		switch c {
		case 'N':
			p.move(north, units)
		case 'S':
			p.move(south, units)
		case 'E':
			p.move(east, units)
		case 'W':
			p.move(west, units)
		case 'L':
			p.turn(left, units)
		case 'R':
			p.turn(right, units)
		case 'F':
			p.forward(units)
		}

		fmt.Println(p.pos)

	})

	fmt.Println(p.manhattan())

}
