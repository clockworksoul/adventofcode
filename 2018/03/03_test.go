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
