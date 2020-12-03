package adventofcode

import (
	"fmt"
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

func TestIngestFileE(t *testing.T) {
	expected := []string{"one", "two", "three"}
	lines := []string{}

	err := IngestFileE("./input_test.txt", func(txt string) error {
		lines = append(lines, txt)
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(expected), len(lines))
	assert.EqualValues(t, expected, lines)
}

func TestIngestFileENoSuchFile(t *testing.T) {
	err := IngestFileE("./foo", func(txt string) error {
		return nil
	})
	assert.NotNil(t, err)
}

func TestIngestFileEWithError(t *testing.T) {
	err := IngestFileE("./input_test.txt", func(txt string) error {
		return fmt.Errorf("expected error")
	})
	assert.NotNil(t, err)
}
