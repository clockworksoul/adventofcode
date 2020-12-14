package main

import (
	"fmt"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

const input = "input.txt"

const bits = 36

type memory map[int64]int64

func (m memory) sum() int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

type bitmask struct {
	ones   int64
	zeroes int64
	xes    int64
}

func newBitmask(mask string) bitmask {
	var ones, zeroes, xes int64

	for _, c := range mask {
		ones <<= 1
		zeroes <<= 1
		xes <<= 1

		switch c {
		case '1':
			ones |= 1
		case '0':
			zeroes |= 1
		case 'X':
			xes |= 1
		}
	}

	return bitmask{ones, zeroes, xes}
}

func (b bitmask) mask(v int64, countZeroes bool) int64 {
	v |= b.ones
	if countZeroes {
		v &= ^b.zeroes
	}
	return v
}

func (b bitmask) xmasks() map[int64]bool {
	masks := map[int64]bool{0: true}

	for i := 0; i <= bits; i++ {
		var m int64 = 1 << i
		if b.xes&m == m {
			for mask := range masks {
				masks[mask|m] = true
			}
		}
	}

	return masks
}

func (b bitmask) maskAddress(addr int64) map[int64]bool {
	masks := b.xmasks()
	addrs := map[int64]bool{}

	masked := b.mask(addr, false) & ^b.xes

	for m := range masks {
		addrs[masked|m] = true
	}

	return addrs
}

func parseMemoryLine(line string) (int64, int64) {
	parts := strings.Split(line, " = ")

	s := strings.Replace(parts[0], "mem[", "", 1)
	s = strings.Replace(s, "]", "", 1)
	mem := adventofcode.MustParseInt64(s)
	value := adventofcode.MustParseInt64(parts[1])
	return mem, value
}

func main() {
	memory1, memory2 := memory{}, memory{}
	mask := bitmask{}

	adventofcode.MustIngestFile(input, func(line string) {
		if line[0:4] == "mask" {
			mask = newBitmask(line[7:])
		} else {
			mem, val := parseMemoryLine(line)

			memory1[mem] = mask.mask(val, true)

			addrs := mask.maskAddress(mem)
			for addr := range addrs {
				memory2[addr] = val
			}
		}
	})

	fmt.Println("Star One:", memory1.sum())
	fmt.Println("Star Two:", memory2.sum())
}
