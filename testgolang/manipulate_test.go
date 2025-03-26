package testgolang_test

import (
	"neversitup-assignment/testgolang"
	"testing"
)

func Test_Shuffle(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "one char",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:  "two chars",
			input: "ab",
			expected: []string{
				"ab",
				"ba",
			},
		},
		{
			name:  "three chars",
			input: "abc",
			expected: []string{
				"abc",
				"acb",
				"bac",
				"bca",
				"cab",
				"cba",
			},
		},
		{
			name:  "four chars",
			input: "abcd",
			expected: []string{
				"abcd",
				"abdc",
				"acbd",
				"acdb",
				"adbc",
				"adcb",
				"bacd",
				"badc",
				"bcad",
				"bcda",
				"bdac",
				"bdca",
				"cabd",
				"cadb",
				"cbad",
				"cbda",
				"cdab",
				"cdba",
				"dabc",
				"dacb",
				"dbac",
				"dbca",
				"dcab",
				"dcba",
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		result := testgolang.Shuffle(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf(
				"%s length error: expected %d but got %d",
				tc.name,
				len(tc.expected),
				len(result),
			)
		}

		for j := range result {
			if result[j] != tc.expected[j] {
				t.Errorf(
					"%s error: expected %d but got %d",
					tc.name,
					len(tc.expected),
					len(result),
				)
			}
		}
	}
}
