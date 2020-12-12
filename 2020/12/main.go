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
	units := (facing * (degrees / 90)) % 4
	if units < 0 {
		units = (4 + units) % 4
	}

	switch units {
	case 0: // No-op
	case 1:
		p.pos[0], p.pos[1] = p.pos[1], -1*p.pos[0]
	case 2:
		p.pos[0], p.pos[1] = -1*p.pos[0], -1*p.pos[1]
	case 3:
		p.pos[0], p.pos[1] = -1*p.pos[1], p.pos[0]
	}
}

func (p *position) forward(waypoint *position, units int) {
	for i := 0; i < units; i++ {
		p.pos[0] += waypoint.pos[0]
		p.pos[1] += waypoint.pos[1]
	}
}

func (p *position) manhattan() int {
	m := math.Abs(float64(p.pos[0])) + math.Abs(float64(p.pos[1]))
	return int(m)
}

func main() {
	p := &position{}
	waypoint := &position{pos: [2]int{10, 1}}

	adventofcode.IngestFile("input.txt", func(line string) {
		fmt.Print(line, " ")
		c := line[0]
		units, _ := strconv.Atoi(line[1:])

		switch c {
		case 'N':
			waypoint.move(north, units)
		case 'S':
			waypoint.move(south, units)
		case 'E':
			waypoint.move(east, units)
		case 'W':
			waypoint.move(west, units)
		case 'L':
			waypoint.turn(left, units)
		case 'R':
			waypoint.turn(right, units)
		case 'F':
			p.forward(waypoint, units)
		}

		fmt.Println(p.pos, waypoint.pos)

	})

	fmt.Println(p.manhattan())
}
