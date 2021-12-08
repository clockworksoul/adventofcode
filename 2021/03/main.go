package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/clockworksoul/adventofcode"
)

const zero byte = '0'
const one byte = '1'

func main() {
	var inputs []string
	adventofcode.IngestFile("input.txt", func(s string) {
		inputs = append(inputs, s)
	})

	part1(inputs)
	part2(inputs)
}

func part1(inputs []string) {
	var bits = make([]byte, len(inputs[0]))

	for i := range bits {
		var counts = make(map[byte]int)

		for _, line := range inputs {
			counts[line[i]]++
		}

		bits[i] = mostByte(counts)
	}

	gamma := bytesToInt(bits)
	mask := int(math.Pow(2, float64(len(bits))) - 1)
	epsilon := gamma ^ mask
	consumption := gamma * epsilon

	fmt.Printf("gamma=%d epsilon=%d epsilon=%d\n", gamma, epsilon, consumption)
}

func part2(inputs []string) {
	ogen := bitCriteria(inputs, true, one)
	co2s := bitCriteria(inputs, false, zero)

	fmt.Printf("ogen=%d co2s=%d rating=%d\n", ogen, co2s, ogen*co2s)
}

func bytesToInt(bits []byte) int {
	i, _ := strconv.ParseInt(string(bits), 2, 32)
	return int(i)
}

func mostByte(m map[byte]int) byte {
	if m[zero] <= m[one] {
		return one
	}
	return zero
}

func bitCriteria(inputs []string, most bool, favor byte) int {
	nbits := len(inputs[0])
	var accum = inputs

	for n := 0; n < nbits && len(accum) > 1; n++ {
		var zeroes, ones []string

		for _, input := range accum {
			switch input[n] {
			case '0':
				zeroes = append(zeroes, input)
			case '1':
				ones = append(ones, input)
			default:
				panic(input)
			}
		}

		if most {
			if len(zeroes) > len(ones) {
				accum = zeroes
			} else if len(ones) > len(zeroes) {
				accum = ones
			} else if favor == zero {
				accum = zeroes
			} else if favor == one {
				accum = ones
			} else {
				panic(accum)
			}
		} else {
			if len(zeroes) < len(ones) {
				accum = zeroes
			} else if len(ones) < len(zeroes) {
				accum = ones
			} else if favor == zero {
				accum = zeroes
			} else if favor == one {
				accum = ones
			} else {
				panic(accum)
			}
		}
	}

	i, _ := strconv.ParseInt(accum[0], 2, 32)
	return int(i)
}

// func leastByte(m map[byte]int) byte {
// 	if m[zero] >= m[one] {
// 		return one
// 	}
// 	return zero
// }
