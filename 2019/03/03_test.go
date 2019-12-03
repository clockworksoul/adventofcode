package main

import (
	"sort"
	"testing"
)

func TestParseMovement(t *testing.T) {
	ine := []string{"R8", "U5", "L5", "D3"}
	me := []Movement{
		Movement{DirRight, 8},
		Movement{DirUp, 5},
		Movement{DirLeft, 5},
		Movement{DirDown, 3},
	}

	for i, in := range ine {
		m := ParseMovement(in)
		e := me[i]

		if m.direction != e.direction {
			t.Errorf("direction mismatch: expected=%v, got=%v\n",
				e.direction, m.direction)
		}

		if m.distance != e.distance {
			t.Errorf("distance mismatch: expected=%d, got=%d\n",
				e.distance, m.distance)
		}
	}
}

func TestTranslate(t *testing.T) {
	ine := []string{"L1", "R1", "U1", "D1"}
	ce := []Coordinate{
		Coordinate{Origin.x - 1, Origin.y},
		Coordinate{Origin.x + 1, Origin.y},
		Coordinate{Origin.x, Origin.y + 1},
		Coordinate{Origin.x, Origin.y - 1},
	}

	for i, in := range ine {
		e := ce[i]
		m := ParseMovement(in)
		c := Origin.Translate(m)

		if c.x != e.x {
			t.Errorf("x mismatch: expected=%v, got=%v\n",
				e.x, c.x)
		}

		if c.y != e.y {
			t.Errorf("y mismatch: expected=%d, got=%d\n",
				e.y, c.y)
		}
	}
}

func TestHashcode(t *testing.T) {
	const MinCoord = -1024
	const MaxCoord = 1024

	m := make(map[string]bool)

	for x := 1; x < MaxCoord; x <<= 1 {
		for y := 1; y < MaxCoord; y <<= 1 {
			c := Coordinate{x, y}
			hash := c.Hashcode()

			if m[hash] {
				t.Errorf("hash collision: %s\n", hash)
			}

			m[hash] = true
		}
	}

	for x := MinCoord; x > 0; x = x >> 1 {
		for y := MinCoord; y > 0; y = y >> 1 {
			c := Coordinate{x, y}
			hash := c.Hashcode()

			if m[hash] {
				t.Errorf("hash collision: %s\n", hash)
			}

			m[hash] = true
		}
	}
}

func TestDistance(t *testing.T) {
	ine := []string{"L1", "R1", "U1", "D1"}

	for _, in := range ine {
		m := ParseMovement(in)
		c := Origin.Translate(m)

		d1 := Distance(Origin, c)
		if d1 != 1 {
			t.Errorf("d1 mismatch: expected=%d, got=%d\n", 1, d1)
		}

		d2 := Distance(c, Origin)
		if d2 != 1 {
			t.Errorf("d2 mismatch: expected=%d, got=%d\n", 1, d2)
		}

		d3 := Distance(c, c)
		if d3 != 0 {
			t.Errorf("d3 mismatch: expected=%d, got=%d\n", 0, d3)
		}
	}
}

func TestParseLine(t *testing.T) {
	line := "R8,U5,L5,D3"

	me := []Movement{
		Movement{DirRight, 8},
		Movement{DirUp, 5},
		Movement{DirLeft, 5},
		Movement{DirDown, 3},
	}

	mm := ParseLine(line)

	for i, m := range mm {
		e := me[i]

		if m.direction != e.direction {
			t.Errorf("direction mismatch: expected=%v, got=%v\n",
				e.direction, m.direction)
		}

		if m.distance != e.distance {
			t.Errorf("distance mismatch: expected=%d, got=%d\n",
				e.distance, m.distance)
		}
	}
}

func TestTraceRoute(t *testing.T) {
	var traces []Trace

	traces = TraceRoute(Coordinate{0, 0}, Coordinate{0, 0})
	if len(traces) != 1 {
		t.Error("unexpected length")
	}
	if traces[0].x != 0 || traces[0].y != 0 {
		t.Errorf("unexpected coordinate: %v\n", traces[0])
	}

	traces = TraceRoute(Coordinate{0, 0}, Coordinate{2, 0})
	if len(traces) != 3 {
		t.Error("unexpected length")
	}
	t.Log(traces)

	traces = TraceRoute(Coordinate{0, 0}, Coordinate{0, 2})
	if len(traces) != 3 {
		t.Error("unexpected length")
	}

	t.Log(traces)
}

func TestFindIntersections(t *testing.T) {
	line1 := "R2,U5"
	line2 := "U2,R5"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	overlaps := FindIntersections(cc1, cc2)

	if len(overlaps) != 1 {
		t.Error("expected 1 overlap")
		return
	}

	o := overlaps[0]
	if o.x != Origin.x+1 && o.y != Origin.y+2 {
		t.Errorf("bad overlap coordinates: %v\n", o)
	}
}

func TestFindIntersections2(t *testing.T) {
	line1 := "R8,U5,L5,D3"
	line2 := "U7,R6,D4,L4"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	overlaps := FindIntersections(cc1, cc2)

	if len(overlaps) != 2 {
		t.Error("expected 2 overlaps")
		return
	}

	t.Log(overlaps)
}

func TestGetAllOccupiedPoints(t *testing.T) {
	line1 := "R8,U5,L5,D3"
	mm1 := ParseLine(line1)
	cc1 := MovementsToCoordinates(mm1)

	points := GetAllOccupiedPoints(cc1)

	for i, p := range points {
		if p.steps != i + 1 {
			t.Errorf("unexpected steps: expected=%d; got=%d\n", i + 1, p.steps)
		}
	}

	t.Log(points)
}

func TestFindClosestIntersection1(t *testing.T) {
	line1 := "R8,U5,L5,D3"
	line2 := "U7,R6,D4,L4"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	t.Log(FindIntersections(cc1, cc2))
	closest := FindClosestIntersection(cc1, cc2)

	t.Log("Coord:   ", closest)
	t.Log("Distance:", Distance(closest.Coordinate, Origin))
}

func TestFindClosestIntersection2(t *testing.T) {
	line1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	line2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	t.Log(FindIntersections(cc1, cc2))
	closest := FindClosestIntersection(cc1, cc2)

	t.Log("Coord:   ", closest)
	t.Log("Distance:", Distance(closest.Coordinate, Origin))
}

func TestFindClosestIntersection3(t *testing.T) {
	line1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	line2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	closest := FindClosestIntersection(cc1, cc2)

	t.Log("Coord:   ", closest)
	t.Log("Distance:", Distance(closest.Coordinate, Origin))
}

func TestSteps1(t *testing.T) {
	line1 := "R8,U5,L5,D3"
	line2 := "U7,R6,D4,L4"
	mm1 := ParseLine(line1)
	mm2 := ParseLine(line2)
	cc1 := MovementsToCoordinates(mm1)
	cc2 := MovementsToCoordinates(mm2)

	ii := FindIntersections(cc1, cc2)

	if len(ii) != 2 {
		t.Error("expected 2 intersections")
		return
	}

	sort.Slice(ii, func(i, j int) bool { return ii[i].Steps() < ii[j].Steps() })

	if ii[0].Steps() != 30 {
		t.Errorf("unexpected steps[0]: expected=%d; got=%d\n", 30, ii[0].Steps())
	}
	if ii[1].Steps() != 40 {
		t.Errorf("unexpected steps[1]: expected=%d; got=%d\n", 40, ii[1].Steps())
	}
	
	t.Log(ii[0])
}
