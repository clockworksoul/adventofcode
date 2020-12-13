package adventofcode

import (
	"bufio"
	"os"
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
