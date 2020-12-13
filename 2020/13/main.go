package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

const input = "input.txt"

func main() {
	starOne()
	starTwo()
}

func starOne() {
	timestamp := 0
	ids := []int{}

	adventofcode.MustIngestFile(input, func(line string) {
		if timestamp == 0 {
			timestamp, _ = strconv.Atoi(line)
		} else {
			parts := strings.Split(line, ",")
			for _, s := range parts {
				if s != "x" {
					id, _ := strconv.Atoi(s)
					ids = append(ids, id)
				}
			}
		}
	})

	eid := 0
	min := math.MaxInt64
	for _, id := range ids {
		time := 0
		for ; time < timestamp; time += id {
		}
		if time < min {
			eid = id
			min = time
		}
	}
	wait := min - timestamp

	fmt.Printf("Star one: %d * %d = %d\n", eid, wait, eid*wait)
}

// Part 2 uses https://en.wikipedia.org/wiki/Chinese_remainder_theorem
// I confess I needed assistance from Reddit on Part 2.
// How can we be expected to just know this?

func starTwo() {
	indices := []int{}
	buses := map[int]int{}

	adventofcode.MustIngestFile(input, func(line string) {
		if len(line) < 10 {
			return
		}

		for i, s := range strings.Split(line, ",") {
			if s == "x" {
				continue
			}

			buses[i] = adventofcode.MustParseInt(s)
			if i == 0 {
				continue
			}
			indices = append(indices, i)
		}
	})

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
