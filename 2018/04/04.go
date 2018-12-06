package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type RecordType int

const (
	RecordTypeUnknown = iota
	RecordTypeBeginsShift
	RecordTypeFallsAsleep
	RecordTypeWakes

	TimeFormat = "2006-01-02 15:04"
)

type Record struct {
	time       time.Time
	recordType RecordType
	guardId    int
}

func (r *Record) String() string {
	var timeString, typeString string

	if r.guardId != 0 {
		typeString = fmt.Sprintf("Guard #%d ", r.guardId)
	}

	switch r.recordType {
	case RecordTypeUnknown:
		typeString += "??"
	case RecordTypeBeginsShift:
		typeString += "begins shift"
	case RecordTypeFallsAsleep:
		typeString += "falls asleep"
	case RecordTypeWakes:
		typeString += "wakes up"
	}

	timeString = r.time.Format(TimeFormat)

	return fmt.Sprintf("[%s] %s", timeString, typeString)
}

// Assumptions:
// 1. Guards fall asleep before they wake up
// 2. Guards always wake up
// 3. Guards begin shift before falling asleep
func CalculateGuardMinutesAsleep(records []*Record) map[int]int {
	sleepMap := make(map[int]int)

	currentGuard := 0
	fellAsleepMinute := 0

	for _, r := range records {
		switch r.recordType {
		case RecordTypeBeginsShift:
			currentGuard = r.guardId
		case RecordTypeFallsAsleep:
			fellAsleepMinute = r.time.Minute()
		case RecordTypeWakes:
			sleepMap[currentGuard] += r.time.Minute() - fellAsleepMinute
		}
	}

	return sleepMap
}

// FindGuardMostAsleepMinute finds the minute at which a guard is most frequently asleep.
// Returns (minute, count)
// Assumptions:
// 1. Records list is sorted
func FindGuardMostAsleepMinute(records []*Record, guardId int) (int, int) {
	currentGuard := 0
	minutes := make([]int, 60)

	fellAsleepMinute, wokeUpMinute := 0, 0

	for _, r := range records {
		switch r.recordType {
		case RecordTypeBeginsShift:
			currentGuard = r.guardId
		case RecordTypeFallsAsleep:
			fellAsleepMinute = r.time.Minute()
		case RecordTypeWakes:
			wokeUpMinute = r.time.Minute()

			if currentGuard == guardId {
				for i := fellAsleepMinute; i < wokeUpMinute; i++ {
					minutes[i]++
				}
			}
		}
	}

	// Find the index with the highest count
	maxMinute, maxCount := 0, 0
	for minute, count := range minutes {
		if count > maxCount {
			maxMinute = minute
			maxCount = count
		}
	}

	return maxMinute, maxCount
}

// Returns (guardId, minute)
// This is brute force. I don't care.
func FindGuardMostFrequentlyAsleepOnSameMinute(records []*Record) (int, int) {
	// Find all guard Ids
	allGuardIds := make(map[int]bool) // Ghetto set

	for _, record := range records {
		if record.recordType == RecordTypeBeginsShift {
			allGuardIds[record.guardId] = true
		}
	}

	maxId, maxMinute, maxCount := 0, 0, 0

	for id, _ := range allGuardIds {
		minute, count := FindGuardMostAsleepMinute(records, id)

		if count > maxCount {
			maxCount = count
			maxId = id
			maxMinute = minute
		}
	}

	return maxId, maxMinute
}

// GetMostSleepingGuard returns (guardId, minutes)
func GetMostSleepingGuard(sleepMap map[int]int) (int, int) {
	maxMins := 0
	maxId := 0

	for id, minutes := range sleepMap {
		if minutes > maxMins {
			maxMins = minutes
			maxId = id
		}
	}

	return maxId, maxMins
}

func IngestRecords(filename string) ([]*Record, error) {
	records := make([]*Record, 0)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record, err := ParseRecord(scanner.Text())
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	SortRecords(records)

	lastId := 0
	for _, r := range records {
		if r.recordType == RecordTypeBeginsShift {
			lastId = r.guardId
		} else {
			r.guardId = lastId
		}
	}

	return records, nil
}

func ParseRecord(value string) (*Record, error) {
	time, text, err := SplitRecordLine(value)
	if err != nil {
		return &Record{}, err
	}

	record := &Record{}

	record.time, err = ParseTime(time)
	if err != nil {
		return record, err
	}

	record.recordType, record.guardId, err = ParseRecordEvent(text)
	if err != nil {
		return record, err
	}

	return record, nil
}

func ParseTime(value string) (time.Time, error) {
	return time.Parse(TimeFormat, value)
}

func ParseRecordEvent(value string) (RecordType, int, error) {
	if len(value) < 5 {
		return RecordTypeUnknown, 0, errors.New("Failed to parse")
	}

	switch value[0:5] {
	case "falls":
		return RecordTypeFallsAsleep, 0, nil
	case "wakes":
		return RecordTypeWakes, 0, nil
	case "Guard":
		var guardId int
		_, err := fmt.Sscanf(value, "Guard #%d begins shift", &guardId)
		return RecordTypeBeginsShift, guardId, err
	default:
		return RecordTypeUnknown, 0, errors.New("Failed to parse")
	}
}

func SortRecords(records []*Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].time.Before(records[j].time)
	})
}

func SplitRecordLine(value string) (string, string, error) {
	split := strings.Split(value, "] ")
	if len(split) != 2 {
		return "", "", errors.New("Invalid string: " + value)
	}

	return split[0][1:], split[1], nil
}

func main() {
	records, err := IngestRecords("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sleepMap := CalculateGuardMinutesAsleep(records)
	guardId, _ := GetMostSleepingGuard(sleepMap)
	minute, count := FindGuardMostAsleepMinute(records, guardId)

	fmt.Printf("Guard %d slept the most on minute %d (%d days) [Answer=%d]\n",
		guardId, minute, count, guardId*minute,
	)

	guardId, minute = FindGuardMostFrequentlyAsleepOnSameMinute(records)
	fmt.Printf("Guard %d is most frequently asleep on the same minute (%d) [Answer=%d]\n",
		guardId, minute, guardId*minute,
	)
}
