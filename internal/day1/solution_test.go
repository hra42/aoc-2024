package day1

import (
	"testing"
)

func TestSolutionPart1(t *testing.T) {
	tests := []struct {
		leftList  []int
		rightList []int
		expected  int
	}{
		{
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			11,
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			0,
		},
		{
			[]int{5, 1, 3},
			[]int{2, 4, 6},
			3,
		},
		{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			0,
		},
	}

	for _, test := range tests {
		result := solutionPart1(test.leftList, test.rightList)
		if result != test.expected {
			t.Errorf("Part 1: For leftList %v and rightList %v, expected %d but got %d",
				test.leftList, test.rightList, test.expected, result)
		}
	}
}

func TestSolutionPart2(t *testing.T) {
	tests := []struct {
		leftList  []int
		rightList []int
		expected  int
	}{
		{
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			31, // 9 + 4 + 0 + 0 + 9 + 9
		},
		{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			9, // Each 1 appears 3 times in right list, so 1*3 + 1*3 + 1*3
		},
		{
			[]int{2, 3, 4},
			[]int{5, 6, 7},
			0, // No numbers in common
		},
		{
			[]int{1, 2, 3},
			[]int{1, 1, 1},
			3, // Only 3 appears in right list (3 times), so 1*3 + 2*0 + 3*0
		},
	}

	for _, test := range tests {
		result := solutionPart2(test.leftList, test.rightList)
		if result != test.expected {
			t.Errorf("Part 2: For leftList %v and rightList %v, expected %d but got %d",
				test.leftList, test.rightList, test.expected, result)
		}
	}
}
