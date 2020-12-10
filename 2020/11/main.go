package main

import (
	"log"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	if err := adventofcode.IngestFile("input.txt", func(line string) {

	}); err != nil {
		log.Fatal(err)
	}
}
