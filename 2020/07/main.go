package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/clockworksoul/adventofcode"
)

type Bag struct {
	Color       string
	Contains    map[string]int
	ContainedBy map[string]int
}

func (b *Bag) ContainsSum() int {
	sum := 1

	for c, i := range b.Contains {
		ib := GetBag(c)
		sum += i * ib.ContainsSum()
	}

	return sum
}

var Bags map[string]*Bag = map[string]*Bag{}

func GetBag(color string) *Bag {
	if bag := Bags[color]; bag == nil {
		Bags[color] = &Bag{Color: color, Contains: map[string]int{}, ContainedBy: map[string]int{}}
		return Bags[color]
	} else {
		return bag
	}
}

func GetAllContainedBy(color string) map[string]*Bag {
	bags := make(map[string]*Bag)
	bag := GetBag(color)

	for c, _ := range bag.ContainedBy {
		bags[c] = GetBag(c)

		for c, b := range GetAllContainedBy(c) {
			bags[c] = b
		}
	}

	return bags
}

func main() {
	p := regexp.MustCompile(` bags?\.?`)

	adventofcode.IngestFile("./input.txt", func(line string) {
		split := strings.Split(line, " bags contain ")
		color := split[0]
		allBags := string(p.ReplaceAll([]byte(split[1]), []byte("")))
		bagSplit := strings.Split(allBags, ", ")

		for _, b := range bagSplit {
			num, p1, p2 := 0, "", ""
			fmt.Sscanf(b, "%d %s %s", &num, &p1, &p2)
			if p1 != "" {
				GetBag(color).Contains[p1+" "+p2] = num
				GetBag(p1 + " " + p2).ContainedBy[color] = num
			}
		}
	})

	fmt.Println(len(GetAllContainedBy("shiny gold")))

	fmt.Println(GetBag("shiny gold").ContainsSum() - 1)
}
