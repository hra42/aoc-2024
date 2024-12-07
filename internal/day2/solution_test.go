package day2

import (
	"testing"
)

func TestIsValidSequence(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for i, test := range tests {
		result := isValidSequence(test.input)
		if result != test.expected {
			t.Errorf("Test case %d: For input %v, expected %v but got %v",
				i, test.input, test.expected, result)
		}
	}
}

func TestIsValidSequenceWithDampener(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},  // Already safe
		{[]int{1, 2, 7, 8, 9}, false}, // Can't be made safe
		{[]int{9, 7, 6, 2, 1}, false}, // Can't be made safe
		{[]int{1, 3, 2, 4, 5}, true},  // Safe by removing 3
		{[]int{8, 6, 4, 4, 1}, true},  // Safe by removing one 4
		{[]int{1, 3, 6, 7, 9}, true},  // Already safe
	}

	for i, test := range tests {
		result := isValidSequenceWithDampener(test.input)
		if result != test.expected {
			t.Errorf("Test case %d: For input %v, expected %v but got %v",
				i, test.input, test.expected, result)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"7 6 4 2 1", []int{7, 6, 4, 2, 1}},
		{"1 2 7 8 9", []int{1, 2, 7, 8, 9}},
	}

	for i, test := range tests {
		result, err := parseLine(test.input)
		if err != nil {
			t.Errorf("Test case %d: Unexpected error: %v", i, err)
			continue
		}

		if len(result) != len(test.expected) {
			t.Errorf("Test case %d: Expected length %d, got %d",
				i, len(test.expected), len(result))
			continue
		}

		for j := range result {
			if result[j] != test.expected[j] {
				t.Errorf("Test case %d: At position %d, expected %d, got %d",
					i, j, test.expected[j], result[j])
			}
		}
	}
}
