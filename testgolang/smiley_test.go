package testgolang_test

import (
	"neversitup-assignment/testgolang"
	"testing"
)

func Test_Smiley(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty input",
			input:    []string{},
			expected: 0,
		},
		{
			name:     "one smiley",
			input:    []string{":)"},
			expected: 1,
		},
		{
			name: "no smiley",
			input: []string{
				":(",
				":I",
				"not smiley emoji",
			},
			expected: 0,
		},
		{
			name: "all smiley",
			input: []string{
				":)",
				":D",
				":-)",
				":~D",
				";D",
				";-D",
				";~)",
			},
			expected: 7,
		},
		{
			name: "many smiley with not smiley",
			input: []string{
				":)",
				":D",
				":-(",
				":~D",
				";D",
				";-P",
				";~)",
			},
			expected: 5,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		result := testgolang.CountSmileys(tc.input)
		if result != tc.expected {
			t.Errorf(
				"%s error: expected %d but got %d",
				tc.name,
				tc.expected,
				result,
			)
		}
	}
}
