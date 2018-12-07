package main

import (
	"testing"
)

func TestIngestPolymer(t *testing.T) {
	polymer, err := IngestPolymer("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	if polymer != "dabAcCaCBAcCcaDA" {
		t.Error("Mismatch")
	}
}

func TestCollapsePolymer(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	collapsed, _ := CollapsePolymer(polymer)

	if collapsed != "dabCBAcaDA" {
		t.Error("Mismatch")
	}
}
