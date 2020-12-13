package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	timestamp := 0
	ids := []int{}

	indices := []int{}
	buses := map[int]int{}

	adventofcode.MustIngestFile("input.txt", func(line string) {
		if timestamp == 0 {
			timestamp = adventofcode.MustParseInt(line)
		} else {
			parts := strings.Split(line, ",")
			for i, s := range parts {
				if s == "x" {
					continue
				}

				si := adventofcode.MustParseInt(s)
				ids = append(ids, si)
				buses[i] = si

				if i != 0 {
					indices = append(indices, i)
				}
			}
		}
	})

	starOne(timestamp, ids)
	starTwo(indices, buses)
}

func starOne(timestamp int, ids []int) {
	eid, min := 0, math.MaxInt64

	for _, id := range ids {
		time := 0
		for ; time < timestamp; time += id {
		}
		if time < min {
			eid = id
			min = time
		}
	}

	fmt.Printf("Star one: %d\n", eid*(min-timestamp))
}

// Part 2 uses https://en.wikipedia.org/wiki/Chinese_remainder_theorem
// I confess I needed assistance from Reddit on Part 2.
// How can we be expected to just know this?

func starTwo(indices []int, buses map[int]int) {
	t := 0
	n1 := buses[0]

	for _, i := range indices {
		var j int

		for {
			j++
			jin := t + j*n1

			test := (jin + i) % buses[i]
			if test != 0 {
				continue
			}

			t = jin
			n1 = n1 * buses[i]
			break
		}
	}

	fmt.Printf("Star two: %d\n", t)
}
