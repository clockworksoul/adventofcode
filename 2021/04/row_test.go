package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowDraw(t *testing.T) {
	tests := []struct {
		numbers  []int
		draw     int
		expected bool
	}{
		{
			numbers:  []int{14, 21, 17, 24, 4},
			draw:     0,
			expected: false,
		},
		{
			numbers:  []int{14, 21, 17, 24, 4},
			draw:     14,
			expected: true,
		},
	}

	for _, test := range tests {
		r := &Row{Numbers: test.numbers}
		result := r.Draw(test.draw)
		assert.Equal(t, test.expected, result)
	}
}

func TestRowHas(t *testing.T) {
	tests := []struct {
		numbers  []int
		n        int
		expected int
	}{
		{
			numbers:  []int{14, 21, 17, 24, 4},
			n:        0,
			expected: -1,
		},
		{
			numbers:  []int{14, 21, 17, 24, 4},
			n:        14,
			expected: 0,
		},
		{
			numbers:  []int{14, 21, 17, 24, 4},
			n:        4,
			expected: 4,
		},
	}

	for _, test := range tests {
		r := &Row{Numbers: test.numbers}
		result := r.Has(test.n)
		assert.Equal(t, test.expected, result)
	}
}

func TestRowWinner(t *testing.T) {
	tests := []struct {
		numbers  []int
		draws    []int
		expected bool
	}{
		{
			numbers:  []int{1},
			draws:    []int{},
			expected: false,
		},
		{
			numbers:  []int{14, 21, 17, 24, 4},
			draws:    []int{14, 21, 17, 24},
			expected: false,
		},
		{
			numbers:  []int{1},
			draws:    []int{1},
			expected: true,
		},
		{
			numbers:  []int{14, 21, 17, 24, 4},
			draws:    []int{14, 21, 17, 24, 4},
			expected: true,
		},
	}

	for _, test := range tests {
		r := &Row{Numbers: test.numbers, Draws: test.draws}
		result := r.Winner()
		assert.Equal(t, test.expected, result)
	}
}
