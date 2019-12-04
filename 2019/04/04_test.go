package main

import (
	"testing"
)

func TestDigits(t *testing.T) {
	num := 278384
	expected := []int{2, 7, 8, 3, 8, 4}
	digits := Digits(num, make([]int, 6))

	if len(digits) != len(expected) {
		t.Fail()
	}

	for i, n := range expected {
		if digits[i] != n {
			t.Fail()
		}
	}

	t.Log(digits)
}

func TestCanBePassword(t *testing.T) {
	tests := map[int]bool{
		122345: true,
		111123: true,
		135679: false,
		111111: true,
		223450: false,
		123789: false,
	}

	for num, expected := range tests {
		digits := Digits(num, make([]int, 6))
		test := CanBePassword(digits)

		if test != expected {
			t.Errorf("%d: expected=%v; got=%v\n", num, expected, test)
		}
	}
}

func TestHasExactlyTwoAdjacent(t *testing.T) {
	tests := map[int]bool{
		112233: true,
		123444: false,
		111122: true,
	}

	for num, expected := range tests {
		digits := Digits(num, make([]int, 6))
		test := HasExactlyTwoAdjacent(digits)

		if test != expected {
			t.Errorf("%d: expected=%v; got=%v\n", num, expected, test)
		}
	}
}
