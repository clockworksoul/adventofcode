package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

func main() {
	valid := 0
	current := []string{}

	adventofcode.IngestFile("./input.txt", func(txt string) {
		if txt == "" {
			if isValid(current) {
				valid++
			}
			current = []string{}
		} else {
			current = append(current, strings.Split(txt, " ")...)
		}
	})
	if isValid(current) {
		valid++
	}

	fmt.Println(valid)
}

func isValid(fields []string) bool {
	expected := map[string]func(string) (bool, error){
		"byr": funcMinMax(1920, 2002),
		"iyr": funcMinMax(2010, 2020),
		"eyr": funcMinMax(2020, 2030),
		"hgt": funcHeight,
		"hcl": func(str string) (bool, error) {
			return regexp.Match("^#[0-9a-f]{6}$", []byte(str))
		},
		"ecl": func(str string) (bool, error) {
			return regexp.Match("^(amb|blu|brn|gry|grn|hzl|oth)$", []byte(str))
		},
		"pid": func(str string) (bool, error) {
			return regexp.Match("^[0-9]{9}$", []byte(str))
		},
		// "cid": true,
	}
	counter := 0

	for _, str := range fields {
		values := strings.Split(str, ":")
		if f, ok := expected[values[0]]; ok {
			if match, err := f(values[1]); err != nil {
				log.Println("warning:", err)
			} else if match {
				counter++
			}
		}
	}

	return counter == len(expected)
}

func funcMinMax(min, max int) func(string) (bool, error) {
	return func(str string) (bool, error) {
		if i, err := strconv.Atoi(str); err != nil {
			return false, err
		} else {
			return min <= i && i <= max, nil
		}
	}
}

func funcHeight(str string) (bool, error) {
	unit := str[len(str)-2:]
	val := str[:len(str)-2]

	switch unit {
	case "cm":
		return funcMinMax(150, 193)(val)
	case "in":
		return funcMinMax(59, 76)(val)
	default:
		return false, nil
	}
}
