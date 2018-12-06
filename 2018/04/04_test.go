package main

import (
	"testing"
)

func TestCalculateGuardMinutesAsleep(t *testing.T) {
	records, err := IngestRecords("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	sleepMap := CalculateGuardMinutesAsleep(records)

	if sleepMap[10] != 50 {
		t.Errorf("Guard 10: expected 50; got %d", sleepMap[10])
	}
}

func TestGetMostSleepingGuard(t *testing.T) {
	records, err := IngestRecords("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	sleepMap := CalculateGuardMinutesAsleep(records)
	id, mins := GetMostSleepingGuard(sleepMap)

	if id != 10 {
		t.Errorf("Id: expected 10; got %d", id)
	}

	if mins != 50 {
		t.Errorf("Mins: expected 50; got %d", mins)
	}
}

func TestFindGuardMostAsleepMinute(t *testing.T) {
	records, err := IngestRecords("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	sleepMap := CalculateGuardMinutesAsleep(records)
	guardId, _ := GetMostSleepingGuard(sleepMap)

	minutes, count := FindGuardMostAsleepMinute(records, guardId)

	if minutes != 24 {
		t.Errorf("Id: expected 24; got %d", minutes)
	}

	if count != 2 {
		t.Errorf("Mins: expected 2; got %d", count)
	}
}

func TestIngestRecords(t *testing.T) {
	records, err := IngestRecords("test.txt")
	if err != nil {
		t.Error(err.Error())
	}

	for i, r := range records {
		if r.guardId == 0 {
			t.Errorf("Line %d: Record has no id", i+1)
		}
	}
}

func TestParseTime(t *testing.T) {
	value := "1518-11-12 23:56"

	time, err := ParseTime(value)
	if err != nil {
		t.Error(err.Error())
	}

	if time.Year() != 1518 {
		t.Error("Year")
	}

	if time.Month() != 11 {
		t.Error("Month")
	}

	if time.Day() != 12 {
		t.Error("Day")
	}

	if time.Hour() != 23 {
		t.Error("Hour")
	}

	if time.Minute() != 56 {
		t.Error("Minute")
	}
}

func TestParseRecordEvent(t *testing.T) {
	rtype, id, err := ParseRecordEvent("Guard #743 begins shift")
	if err != nil || rtype != RecordTypeBeginsShift || id != 743 {
		t.Error("Guard #743 begins shift", rtype, id, err)
	}

	rtype, id, err = ParseRecordEvent("falls asleep")
	if err != nil || rtype != RecordTypeFallsAsleep || id != 0 {
		t.Error("falls asleep", rtype, id, err)
	}

	rtype, id, err = ParseRecordEvent("wakes up")
	if err != nil || rtype != RecordTypeWakes || id != 0 {
		t.Error("wakes up", rtype, id, err)
	}

	rtype, id, err = ParseRecordEvent("blah")
	if err == nil {
		t.Error()
	}
}

func TestParseRecord(t *testing.T) {
	value := "[1518-03-25 00:01] Guard #743 begins shift"

	record, err := ParseRecord(value)
	if err != nil {
		t.Error(err.Error())
	}

	if record.guardId != 743 {
		t.Error("GuardId")
	}

	if record.recordType == RecordTypeUnknown {
		t.Error("Record type unknown")
	}
}
