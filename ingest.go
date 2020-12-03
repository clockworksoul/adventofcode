package adventofcode

import (
	"bufio"
	"os"
)

func IngestFile(name string, f func(string)) error {
	return IngestFileE(name, func(s string) error {
		f(s)
		return nil
	})
}

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

	return nil
}
