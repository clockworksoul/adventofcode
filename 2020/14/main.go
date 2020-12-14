package main

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	memory := map[string]int64{}
	mask := bitmask{}

	adventofcode.MustIngestFile("input.txt", func(line string) {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			mask = newBitmask(parts[1])
		} else {
			val := mask.mask(adventofcode.MustParseInt64(parts[1]))
			memory[parts[0]] = val
			fmt.Println(parts[0], val)
		}
	})

	var sum int64
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

type bitmask struct {
	ones   int64
	zeroes int64
}

func newBitmask(mask string) bitmask {
	var ones, zeroes int64

	for _, c := range mask {
		ones <<= 1
		zeroes <<= 1

		if c == '1' {
			ones |= 1
		} else if c == '0' {
			zeroes |= 1
		}
	}

	return bitmask{ones, zeroes}
}

func (b bitmask) mask(v int64) int64 {
	v |= b.ones
	v &= ^b.zeroes
	return v
}
