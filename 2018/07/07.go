package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Step struct {
	Label  string
	before []*Step
	after  []*Step
}

func (s *Step) RemoveBefore(toRemove *Step) {
	for i, b := range s.before {
		if b.Label == toRemove.Label {
			s.before = append(s.before[:i], s.before[i+1:]...)
			break
		}
	}
}

func (s *Step) String() string {
	return s.Label
}

type Steps map[string]*Step

func (s Steps) Get(key string) *Step {
	if _, ok := s[key]; !ok {
		s[key] = &Step{Label: key}
	}

	return s[key]
}

func IngestSteps(filename string) (Steps, error) {
	steps := Steps{}

	file, err := os.Open(filename)
	if err != nil {
		return steps, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var before, after string

		line := scanner.Text()
		_, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &before, &after)
		if err != nil {
			return steps, err
		}

		beforeStep := steps.Get(before)
		afterStep := steps.Get(after)

		beforeStep.after = append(beforeStep.after, afterStep)
		afterStep.before = append(afterStep.before, beforeStep)
	}

	return steps, nil
}

func FindStartingSteps(steps Steps) []*Step {
	s := make([]*Step, 0)

	for _, step := range steps {
		if len(step.before) == 0 {
			s = append(s, step)
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return strings.Compare(s[i].Label, s[j].Label) < 0
	})

	return s
}

func GetExecutionSteps(steps Steps) []*Step {
	stepsSlice := make([]*Step, 0)

	for len(steps) > 0 {
		nextStep := FindStartingSteps(steps)[0]
		for _, after := range nextStep.after {
			after.RemoveBefore(nextStep)
		}

		stepsSlice = append(stepsSlice, nextStep)
		delete(steps, nextStep.Label)
	}

	return stepsSlice
}

func GetExecutionStepsString(steps Steps) string {
	str := ""

	for _, s := range GetExecutionSteps(steps) {
		str += s.Label
	}

	return str
}

func main() {
	steps, err := IngestSteps("input.txt")
	if err != nil {
		log.Fatal()
	}

	esteps := GetExecutionStepsString(steps)

	log.Println(esteps)
}
