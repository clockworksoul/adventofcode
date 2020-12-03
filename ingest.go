package adventofcode

import (
	"bufio"
	"os"
)

func IngestFile(name string, f func(string)) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}

	return nil
}
