package testgolang_test

import (
	"neversitup-assignment/testgolang"
	"testing"
)

func Test_OddNumber(t *testing.T) {
	testCases := []struct {
		testName string
		input    []int
		expected int
	}{
		{
			testName: "nil input",
			input:    nil,
			expected: 0,
		},
		{
			testName: "empty input",
			input:    []int{},
			expected: 0,
		},
		{
			testName: "distinct numbers",
			input:    []int{1, 2, 3, 4},
			expected: 4,
		},
		{
			testName: "no odd number",
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: 0,
		},
		{
			testName: "one odd number",
			input:    []int{1, 2, 2, 3, 3, 3, 3, 4, 4},
			expected: 1,
		},
		{
			testName: "many odd numbers",
			input:    []int{1, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4},
			expected: 4,
		},
		{
			testName: "many odd numbers, many even numbers",
			input:    []int{1, 2, 2, 3, 3, 3, 3, 3, 4, 4},
			expected: 2,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		result := testgolang.OddNumber(tc.input)
		if result != tc.expected {
			t.Errorf(
				"%s error: expected %d but got %d",
				tc.testName,
				tc.expected,
				result,
			)
		}
	}
}
