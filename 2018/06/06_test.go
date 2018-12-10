package main

import (
	"testing"
)

func TestIngestPoints(t *testing.T) {
	points, err := IngestPoints("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(points)
}

func TestDistanceTo(t *testing.T) {
	points := []*Point{
		&Point{X: 0, Y: 0},
		&Point{X: 1, Y: 0},
		&Point{X: 1, Y: 1},
	}

	if points[0].DistanceTo(points[0]) != 0 {
		t.Error()
	}

	if points[0].DistanceTo(points[1]) != 1 {
		t.Error()
	}

	if points[0].DistanceTo(points[2]) != 2 {
		t.Error()
	}
}

func TestSizeOfClosestAreas(t *testing.T) {
	points, err := IngestPoints("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	m := SizeOfClosestAreas(points)

	margins := FindMargins(points)
	for _, p := range margins {
		if m[*p] != -1 {
			t.Error("Expected infinite area")
		}
	}

	if m[Point{X: 3, Y: 4}] != 9 {
		t.Errorf("Expected 9, got %d", m[Point{X: 3, Y: 4}])
	}

	if m[Point{X: 5, Y: 5}] != 17 {
		t.Errorf("Expected 17, got %d", m[Point{X: 5, Y: 5}])
	}
}

func TestFindMargins(t *testing.T) {
	points, err := IngestPoints("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	margins := FindMargins(points)

	t.Log(margins)

	expected := []*Point{
		&Point{X: 1, Y: 1},
		&Point{X: 8, Y: 3},
		&Point{X: 8, Y: 9},
		&Point{X: 1, Y: 6},
	}

	if len(margins) != len(expected) {
		t.Error()
	}

	for i := range expected {
		if *expected[i] != *margins[i] {
			t.Errorf("%s vs %s\n", expected[i], margins[i])
		}
	}
}

func TestAngleTo(t *testing.T) {
	origin := &Point{X: 0, Y: 0}
	var p *Point

	// By our convention, the angle of a point to itself shall be 360
	p = origin
	if origin.AngleTo(origin) != 360 {
		t.Errorf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: 2, Y: 0}
	if origin.AngleTo(p) != 0 {
		t.Errorf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: 2, Y: 2}
	if origin.AngleTo(p) != 45 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: 0, Y: 2}
	if origin.AngleTo(p) != 90 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: -2, Y: 2}
	if origin.AngleTo(p) != 135 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: -2, Y: 0}
	if origin.AngleTo(p) != 180 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: -2, Y: -2}
	if origin.AngleTo(p) != 225 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: 0, Y: -2}
	if origin.AngleTo(p) != 270 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}

	p = &Point{X: 2, Y: -2}
	if origin.AngleTo(p) != 315 {
		t.Logf("%s: %.02f\n", p, origin.AngleTo(p))
	}
}
