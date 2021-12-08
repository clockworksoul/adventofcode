package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestBoardWinner(t *testing.T) {
	tests := []struct {
		rows     [][]int
		draws    []int
		expected bool
	}{
		{
			rows: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6}},
			draws:    []int{3, 15, 0, 2, 22},
			expected: true,
		},
		{
			rows: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6}},
			draws:    []int{3, 9, 19, 20, 14},
			expected: true,
		},
		{
			rows: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6}},
			draws:    []int{3, 15, 0, 2, 17},
			expected: false,
		},
		{
			rows: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6}},
			draws:    []int{3, 9, 19, 20, 11},
			expected: false,
		},
	}

	for _, test := range tests {
		b := NewBoard(5)

		for _, nn := range test.rows {
			b.AddRow(&Row{Numbers: nn})
		}

		for _, d := range test.draws {
			b.Draw(d)
			t.Log(d, b.Winner())
		}

		assert.Equal(t, test.expected, b.Winner())
		// result := r.Winner()
		// assert.Equal(t, test.expected, result)
	}
}
