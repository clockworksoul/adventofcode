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
	assert.Error(t, err)
}

func TestIngestFileEWithError(t *testing.T) {
	err := IngestFileE("./input_test.txt", func(txt string) error {
		return fmt.Errorf("expected error")
	})
	assert.Error(t, err)
}

func TestParseInts(t *testing.T) {
	in := []string{"1", "2", "3", "4", "99"}
	expected := []int{1, 2, 3, 4, 99}
	out, err := ParseInts(in)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestSplitAndParseInts(t *testing.T) {
	tests := []struct {
		s        string
		regex    string
		expected []int
		err      bool
	}{
		{
			s:        "7,4,9,5,11,17",
			regex:    ",",
			expected: []int{7, 4, 9, 5, 11, 17},
		},
		{
			s:        "22 13 17 11  0",
			regex:    " +",
			expected: []int{22, 13, 17, 11, 0},
		},
		{
			s:     "7,4,9,5,11,foo",
			regex: ",",
			err:   true,
		},
	}

	for _, test := range tests {
		msg := "s=%q regex=%q err=%t"
		ii, err := SplitAndParseInts(test.s, test.regex)

		if test.err {
			assert.Error(t, err, msg, test.s, test.regex, test.err)
			continue
		}

		assert.NoError(t, err, msg, test.s, test.regex, test.err)
		assert.Equal(t, test.expected, ii, msg, test.s, test.regex, test.err)
	}
}
