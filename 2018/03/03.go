package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EncodingBits = 16
	BitMask      = 0xFFFF
)

type Line struct {
	Id                int
	LeftEdge, TopEdge int
	Width, Height     int
}

func DecodePosition(encoded uint64) (int, int) {
	height := int(encoded & BitMask)
	width := int(encoded >> EncodingBits)
	return width, height
}

func EncodePosition(width, height int) uint64 {
	return uint64(width)<<EncodingBits | uint64(height)
}

func LineReader(filename string) (chan Line, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	c := make(chan Line)
	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			line, err := ParseLine(scanner.Text())
			if err != nil {
				panic(err.Error())
			}
			c <- line
		}

		close(c)
	}()

	return c, nil
}

// "#1 @ 1,3: 4x4",
func ParseLine(str string) (Line, error) {
	line := Line{}

	_, err := fmt.Sscanf(str,
		"#%d @ %d,%d: %dx%d",
		&line.Id,
		&line.LeftEdge, &line.TopEdge,
		&line.Width, &line.Height,
	)
	if err != nil {
		return line, err
	}

	return line, nil
}

func BuildOverLapsMap(ch chan Line) map[uint64][]int {
	overlaps := make(map[uint64][]int)

	for line := range ch {
		for x := line.LeftEdge; x < line.LeftEdge+line.Width; x++ {
			for y := line.TopEdge; y < line.TopEdge+line.Height; y++ {
				overlaps[EncodePosition(x, y)] = append(overlaps[EncodePosition(x, y)], line.Id)
			}
		}
	}

	return overlaps
}

func CountOverlaps(ch chan Line) int {
	overlaps := BuildOverLapsMap(ch)
	overLapCount := 0

	for _, v := range overlaps {
		if len(v) >= 2 {
			overLapCount++
		}
	}

	return overLapCount
}

func GetNonoverlappers(ch chan Line) []int {
	overlaps := BuildOverLapsMap(ch)
	overlappers := make(map[int]bool)

	for _, v := range overlaps {
		for _, id := range v {
			if _, exists := overlappers[id]; !exists {
				overlappers[id] = len(v) > 1
			} else if len(v) > 1 {
				overlappers[id] = true
			}
		}
	}

	nonoverlappers := make([]int, 0)

	for id, v := range overlappers {
		if !v {
			nonoverlappers = append(nonoverlappers, id)
		}
	}

	return nonoverlappers
}

func main() {
	// overLapCount := CountOverlaps(ch)
	// fmt.Println(overLapCount)

	ch, err := LineReader("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, id := range GetNonoverlappers(ch) {
		fmt.Println(id)
	}
}
