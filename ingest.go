package adventofcode

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// MustIngestFile ingests the text file name and calls function f on each line.
// The function will panic if a non-EOF error is encountered by the scanner.
func MustIngestFile(name string, f func(string)) {
	if err := IngestFile(name, f); err != nil {
		panic(err)
	}
}

// IngestFile ingests the text file name and calls function f on each line.
// A non-nil error will be returned if a non-EOF error is encountered
// by the scanner.
func IngestFile(name string, f func(string)) error {
	return IngestFileE(name, func(s string) error {
		f(s)
		return nil
	})
}

// MustIngestFileE ingests the text file name and calls function f on each line.
// The function will panic if a non-nil error is f returns an error or if a
// non-EOF error is encountered by the scanner.
func MustIngestFileE(name string, f func(string) error) {
	if err := IngestFileE(name, f); err != nil {
		panic(err)
	}
}

// IngestFileE ingests the text file name and calls function f on each line.
// A non-nil error will be returned if f returns an error, or if a
// non-EOF error is encountered by the scanner.
func IngestFileE(name string, f func(string) error) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = f(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// IngestFileInt ingests the text file name and calls function f on each line.
// It will attempt to convert each line to an integer using strconv.Atoi(s).
// A non-nil error will be returned if a non-EOF error is encountered
// by the scanner or if a line is not convertable into an int.
func IngestFileInt(name string, f func(int)) error {
	return IngestFileE(name, func(s string) error {
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		f(val)
		return nil
	})
}

// IngestFileIntE ingests the text file name and calls function f on each line.
// It will attempt to convert each line to an integer using strconv.Atoi(s).
// A non-nil error will be returned if f returns an error, a non-EOF error is
// encountered by the scanner, or a line is not convertable into an int.
func IngestFileIntE(name string, f func(int) error) error {
	return IngestFileE(name, func(s string) error {
		val, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return f(val)
	})
}

// MustParseInt parses a string into an int or panics.
func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// MustParseInt64 parses a string into an int64 or panics.
func MustParseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// ParseInts parses a slice of strings into a slice of ints.
func ParseInts(ss []string) ([]int, error) {
	var nn = make([]int, len(ss))

	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nn[i] = n
	}
	return nn, nil
}

// MustParseInts parses a slice of strings into a slice of ints or panics.
func MustParseInts(ss []string) []int {
	nn, err := ParseInts(ss)
	if err != nil {
		panic(err)
	}
	return nn
}

// SplitAndParseInts executes a regex split of a string and converts
// its contents into integers.
func SplitAndParseInts(s string, rexp string) ([]int, error) {
	re, err := regexp.Compile(rexp)
	if err != nil {
		return nil, err
	}
	split := re.Split(s, -1)
	nn := make([]int, len(split))

	for i, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		nn[i] = n
	}
	return nn, nil
}

// MustSplitAndParseInts executes a regex split of a string and converts
// its contents into integers, or it panics.
func MustSplitAndParseInts(s string, rexp string) []int {
	nn, err := SplitAndParseInts(s, rexp)
	if err != nil {
		panic(err)
	}
	return nn
}
