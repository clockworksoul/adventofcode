package main

import "fmt"

const (
	RangeStart = 278384
	RangeEnd   = 824795
)

func main() {
	day1, day2 := 0, 0

	digits := make([]int, 6)
	for num := RangeStart; num <= RangeEnd; num++ {
		digits := Digits(num, digits)

		if CanBePassword(digits) {
			day1++

			if HasExactlyTwoAdjacent(digits) {
				day2++
			}
		}
	}

	fmt.Println(day1, day2)
}

// It is a six-digit number.
// The value is within the range given in your puzzle input.
// Two adjacent digits are the same (like 22 in 122345).
// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
func CanBePassword(digits []int) bool {
	adjacent := false
	a := digits[0]

	for i := 1; i < len(digits); i++ {
		b := digits[i]

		if b == a {
			adjacent = true
		}

		if a > b {
			return false
		}

		a = b
	}

	return adjacent
}

func HasExactlyTwoAdjacent(digits []int) bool {
	// We have at least one matching digit.
	// Two passes is hacky and I know it.

	a := digits[0]
	sequence := 1

	for i := 1; i < len(digits); i++ {
		b := digits[i]

		if b == a {
			sequence++
		} else if sequence == 2 {
			return true
		} else {
			sequence = 1
		}

		a = b
	}
	return sequence == 2
}

func Digits(num int, digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = num % 10
		num /= 10
	}

	return digits
}
