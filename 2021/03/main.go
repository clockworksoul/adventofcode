package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	var inputs []string
	adventofcode.IngestFile("input.txt", func(s string) {
		inputs = append(inputs, s)
	})

	part1(inputs)
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
	var mbits = make([]byte, len(inputs[0]))
	var lbits = make([]byte, len(inputs[0]))

	for i := range mbits {
		var counts = make(map[byte]int)

		for _, line := range inputs {
			counts[line[i]]++
		}

		mbits[i] = mostByte(counts)
		lbits[i] = leastByte(counts)
	}
}

func bytesToInt(bits []byte) int {
	i, _ := strconv.ParseInt(string(bits), 2, 32)
	return int(i)
}

const zero byte = '0'
const one byte = '1'

func mostByte(m map[byte]int) byte {
	if m[one] >= m[zero] {
		return one
	}
	return zero
}

func leastByte(m map[byte]int) byte {
	if m[zero] >= m[one] {
		return zero
	}
	return one
}
