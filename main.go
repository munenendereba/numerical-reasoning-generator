package main

import "sort"

func main() {}

func GenerateNumericalReasoningChallenges() {

}

func ReturnFurthestFromMedian(nums []int) int {
	sort.Ints(nums)

	small := nums[0]
	large := nums[2]
	mid := nums[1]

	if large-mid > mid-small {
		return large
	}

	return small
}
