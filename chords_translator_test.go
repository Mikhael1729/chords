package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChordsTranslator_Process(t *testing.T) {
	assert := require.New(t)

	testCases := []struct {
		word               []string
		expectedChordNotes string
	}{
		// {[]string{"B", "min"}, "B, D, F#"},
		// {[]string{"B#"}, "B#, D#, F#"},
		{[]string{"B#", "min"}, "B# D F#"},
	}

	translator := NewChordsTranslator()

	for _, testCase := range testCases {
		notes := translator.Process(testCase.word)
		assert.Equal(testCase.expectedChordNotes, notes)
	}
}
