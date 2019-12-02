package main

import (
	"testing"
)

func TestIngestion(t *testing.T) {
	steps, err := IngestSteps("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	c := steps.Get("C")
	a := steps.Get("A")

	if c.Label != "C" {
		t.Error()
	}

	if a.Label != "A" {
		t.Error()
	}

	if c.after[0] != a {
		t.Error()
	}

	if a.before[0] != c {
		t.Error()
	}
}

func TestGetStartingSteps(t *testing.T) {
	steps, err := IngestSteps("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	starting := FindStartingSteps(steps)

	if starting[0].Label != "C" {
		t.Error()
	}
}

func TestGetExecutionSteps(t *testing.T) {
	steps, err := IngestSteps("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	esteps := GetExecutionStepsString(steps)

	if esteps != "CABDFE" {
		t.Error()
	}
}

func TestRemoveBefore(t *testing.T) {
	steps, err := IngestSteps("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	c := steps.Get("C")

	for _, a := range c.after {
		t.Log("Doing:", *a)
		t.Log(a.before)
		a.RemoveBefore(c)
		t.Log(a.before)
	}

	starting := FindStartingSteps(steps)
	t.Log(starting)
}
