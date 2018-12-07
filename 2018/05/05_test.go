package main

import (
	"testing"
)

func TestCollapsePolymer(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	collapsed, _ := CollapsePolymer(polymer)

	if collapsed != "dabCBAcaDA" {
		t.Error("Mismatch")
	}
}

func TestIngestPolymer(t *testing.T) {
	polymer, err := IngestPolymer("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	if polymer != "dabAcCaCBAcCcaDA" {
		t.Error("Mismatch")
	}
}

func TestFindShortestWithOneUnitRemoved(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"
	length, unit := FindShortestWithOneUnitRemoved(polymer)

	if length != 4 {
		t.Error(length)
	}

	if unit != "C" {
		t.Error(unit)
	}
}

func TestScrubUnits(t *testing.T) {
	polymer := "dabAcCaCBAcCcaDA"

	if ScrubUnits(polymer, "A") != "dbcCCBcCcD" {
		t.Errorf("A")
	}

	if ScrubUnits(polymer, "B") != "daAcCaCAcCcaDA" {
		t.Errorf("B")
	}

	if ScrubUnits(polymer, "C") != "dabAaBAaDA" {
		t.Errorf("C")
	}

	if ScrubUnits(polymer, "D") != "abAcCaCBAcCcaA" {
		t.Errorf("D")
	}
}
