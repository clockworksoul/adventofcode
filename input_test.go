package adventofcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngestFile(t *testing.T) {
	expected := []string{"one", "two", "three"}
	lines := []string{}

	err := IngestFile("./input_test.txt", func(txt string) {
		lines = append(lines, txt)
	})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(expected), len(lines))
	assert.EqualValues(t, expected, lines)
}
