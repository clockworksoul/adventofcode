package main

import (
	"testing"
)

// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.
func TestHasNMultiples(t *testing.T) {
	tests := map[string]int{
		"abcdef": 0,
		"bababc": 5,
		"abbcde": 2,
		"abcccd": 3,
		"aabcdd": 2,
		"abcdee": 2,
		"ababab": 3,
	}

	for k, v := range tests {
		sum := 0
		bools := hasNMultiples(k, 3)

		// Ignore "1" counts
		for i := 2; i < len(bools); i++ {
			if bools[i] {
				sum += i
			}
		}

		if sum != v {
			t.Errorf("%s: expected %d; got %d\n", k, v, sum)
		}
	}
}

func TestDifferenceCount(t *testing.T) {
	var same string
	var diff int

	same, diff = diffStrings("abcde", "axcye")
	if diff != 2 {
		t.Errorf("Expected %d; got %d\n", 2, diff)
	}
	if same != "ace" {
		t.Errorf("Expected %s; got %s\n", "ace", same)
	}

	same, diff = diffStrings("fghij", "fguij")
	if diff != 1 {
		t.Errorf("Expected %d; got %d\n", 1, diff)
	}
	if same != "fgij" {
		t.Errorf("Expected %s; got %s\n", "fgij", same)
	}
}
