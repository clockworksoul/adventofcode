package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		l, err := parseLine(txt)
		if err != nil {
			log.Fatalf("parse failure: \"%s\"", txt)
		}

		if evaluateLine2(l) {
			count++
		}
	}

	fmt.Println(count)
}

func evaluateLine1(l LineParts) bool {
	count := strings.Count(l.password, l.letter)

	return count >= l.min && count <= l.max
}

func evaluateLine2(l LineParts) bool {
	l1 := string(l.password[l.min-1]) == l.letter
	l2 := string(l.password[l.max-1]) == l.letter

	return (l1 || l2) && !(l1 && l2)
}
