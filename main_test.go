package main

import "testing"

var tests = []struct {
	num1   int
	num2   int
	num3   int
	result int
}{
	{1, 2, 4, 4},
	{9, 2, 7, 2},
	{17, 23, 19, 23},
	{8, 5, 3, 8},
}

func TestGenerateNumericalReasoningChallenges(t *testing.T) {
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := ReturnFurthestFromMedian([]int{test.num1, test.num2, test.num3})

			if res != test.result {
				t.Errorf("ReturnFurthestFromMedian(%d,%d,%d) expected %d, but got %d", test.num1, test.num2, test.num3, test.result, res)
			}
		})

	}

}
