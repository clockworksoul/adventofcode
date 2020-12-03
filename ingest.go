package adventofcode

import (
	"bufio"
	"os"
)

// IngestFile ingests the text file name and calls function f on each line.
// A non-nil error will be returned if a non-EOF error is encountered
// by the scanner.
func IngestFile(name string, f func(string)) error {
	return IngestFileE(name, func(s string) error {
		f(s)
		return nil
	})
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
