package main

import (
	"testing"
)

func TestCodec(t *testing.T) {
	width := 3
	height := 5

	encoded := EncodePosition(width, height)
	nw, nh := DecodePosition(encoded)

	if width != nw {
		t.Error("Bad width")
	}

	if height != nh {
		t.Error("Bad height")
	}
}

func TestBuildOverLapsMap(t *testing.T) {
	ch, err := LineReader("test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	overlapsMap := BuildOverLapsMap(ch)

	for encoded, lines := range overlapsMap {
		x, y := DecodePosition(encoded)
		t.Logf("%dx%d --> %v\n", x, y, lines)
	}
}

func TestGetNonoverlappers(t *testing.T) {
	ch, err := LineReader("test.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}

	nonoverlappers := GetNonoverlappers(ch)

	if len(nonoverlappers) != 1 {
		t.Errorf("Expected 1; got %d\n", len(nonoverlappers))
	}

	if nonoverlappers[0] != 3 {
		t.Errorf("Not 3")
	}
}

func TestParseLine(t *testing.T) {
	lineStrings := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	lineSolutions := []Line{
		Line{1, 1, 3, 4, 4},
		Line{2, 3, 1, 4, 4},
		Line{3, 5, 5, 2, 2},
	}

	for i := 0; i < len(lineStrings); i++ {
		line, err := ParseLine(lineStrings[i])
		if err != nil {
			t.Error(err.Error())
		}

		if line != lineSolutions[i] {
			t.Error("Line mismatch")
		}
	}
}
