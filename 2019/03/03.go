package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Direction byte

const (
	_ Direction = iota
	DirUp
	DirDown
	DirLeft
	DirRight
)

// var Origin = Coordinate{1 << 24, 1 << 24}
var Origin = Coordinate{}

type Movement struct {
	direction Direction
	distance  int
}

type Coordinate struct {
	x int
	y int
}

type Trace struct {
	Coordinate
	steps int
}

type Intersection struct {
	Coordinate
	aSteps int
	bSteps int
}

func (i Intersection) Steps() int {
	return i.aSteps + i.bSteps
}

func (c Coordinate) Hashcode() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func (c Coordinate) Translate(m Movement) Coordinate {
	nc := Coordinate{c.x, c.y}

	switch m.direction {
	case DirDown:
		nc.y -= m.distance
	case DirLeft:
		nc.x -= m.distance
	case DirRight:
		nc.x += m.distance
	case DirUp:
		nc.y += m.distance
	}

	return nc
}

func Distance(a, b Coordinate) (d int) {
	if a.x > b.x {
		d += a.x - b.x
	} else {
		d += b.x - a.x
	}

	if a.y > b.y {
		d += a.y - b.y
	} else {
		d += b.y - a.y
	}

	return d
}

func ParseMovement(s string) (m Movement) {
	switch s[0] {
	case 'U':
		m.direction = DirUp
	case 'D':
		m.direction = DirDown
	case 'L':
		m.direction = DirLeft
	case 'R':
		m.direction = DirRight
	}

	fmt.Sscanf(s[1:], "%d", &m.distance)

	return m
}

func ParseLine(str string) []Movement {
	split := strings.Split(str, ",")
	mm := make([]Movement, 0)

	for _, s := range split {
		mm = append(mm, ParseMovement(s))
	}

	return mm
}

func MovementsToCoordinates(mm []Movement) []Coordinate {
	cc := make([]Coordinate, 0)
	c := Origin

	for _, m := range mm {
		c = c.Translate(m)
		cc = append(cc, c)
	}

	return cc
}

func readInputLines() ([]string, error) {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot open input file: %w", err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// TraceRoute assumes a straight line where either X or Y match
func TraceRoute(a, b Coordinate) []Trace {
	traces := make([]Trace, 0)

	// Condition 1: same X coordinates
	if a.x == b.x {
		if a.y < b.y {
			for y := a.y; y <= b.y; y++ {
				traces = append(traces, Trace{Coordinate{a.x, y}, 0})
			}
		} else {
			for y := a.y; y >= b.y; y-- {
				traces = append(traces, Trace{Coordinate{a.x, y}, 0})
			}
		}
	} else if a.y == b.y {
		if a.x < b.x {
			for x := a.x; x <= b.x; x++ {
				traces = append(traces, Trace{Coordinate{x, a.y}, 0})
			}
		} else {
			for x := a.x; x >= b.x; x-- {
				traces = append(traces, Trace{Coordinate{x, a.y}, 0})
			}
		}
	}

	return traces
}

func GetAllOccupiedPoints(ccin []Coordinate) []Trace {
	c := Origin
	traces := []Trace{}

	for i := 0; i < len(ccin); i++ {
		route := TraceRoute(c, ccin[i])
		traces = append(traces, route[1:]...)
		c = ccin[i]
	}

	for i := range traces {
		traces[i].steps = i + 1
	}

	return traces
}

func FindIntersections(ca, cb []Coordinate) []Intersection {
	overlaps := make([]Intersection, 0)
	m := make(map[string][]Intersection)

	aa := GetAllOccupiedPoints(ca)
	bb := GetAllOccupiedPoints(cb)

	for _, ca := range aa {
		i := Intersection{
			Coordinate: ca.Coordinate,
			aSteps:     ca.steps,
		}
		m[ca.Hashcode()] = []Intersection{i}
	}

	for _, cb := range bb {
		ii, ok := m[cb.Hashcode()]

		if ok {
			for _, i := range ii {
				i.bSteps = cb.steps
				overlaps = append(overlaps, i)
			}
		}
	}

	return overlaps
}

func FindClosestIntersection(ca, cb []Coordinate) Intersection {
	intersections := FindIntersections(ca, cb)

	sort.Slice(
		intersections,
		func(i, j int) bool {
			d1 := Distance(Origin, intersections[i].Coordinate)
			d2 := Distance(Origin, intersections[j].Coordinate)
			return d1 < d2
		},
	)

	return intersections[0]
}

func main() {
	lines, err := readInputLines()
	if err != nil {
		log.Fatal(err)
	}

	mm1 := ParseLine(lines[0])
	mm2 := ParseLine(lines[1])

	coords1 := MovementsToCoordinates(mm1)
	coords2 := MovementsToCoordinates(mm2)

	day1(coords1, coords2)
	day2(coords1, coords2)
}

func day1(coords1, coords2 []Coordinate) {
	closest := FindClosestIntersection(coords1, coords2)

	fmt.Println("Day 1:")
	fmt.Println("	Coord:   ", closest)
	fmt.Println("	Distance:", Distance(closest.Coordinate, Origin))
}

func day2(coords1, coords2 []Coordinate) {
	ii := FindIntersections(coords1, coords2)

	sort.Slice(ii, func(i, j int) bool { return ii[i].Steps() < ii[j].Steps() })

	fmt.Println("Day 2:")
	fmt.Println("	Steps:   ", ii[0].Steps())
}
