package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strconv"
)

type num struct {
	x1 int
	x2 int
	x3 int
}

func main() {
	nums := GenerateNumericalReasoningChallenges(20, 5, 4)

	//time.Sleep(4 * 5 * time.Second)

	for _, num := range nums {
		fmt.Println(num.x1, num.x2, num.x3)

		var answer string
		fmt.Scanln(&answer)
		a, _ := strconv.Atoi(answer)

		xx := []int{num.x1, num.x2, num.x3}
		bb := ReturnFurthestFromMedian(xx)

		if a != bb {
			fmt.Printf("***incorrect***, expected %d, but got %d \n", bb, a)
		} else {
			fmt.Printf("!!correct!!, expected %d, and got %d \n", bb, a)
		}
	}
}

func GenerateNumericalReasoningChallenges(maxnum, numquestions, secperq int) []num {
	nums := []num{}

	for x := 1; x <= numquestions; x++ {

		num1 := rand.IntN(maxnum + 1)
		ranger1 := rand.IntN(22)
		ranger2 := rand.IntN(22)
		num2 := num1 + ranger1
		num3 := num2 - ranger2

		//ensure ranger 1 and 2 are not equal
		for ranger2 == ranger1 || num3 < 0 {
			ranger2 = rand.IntN(22)
			num3 = num2 - ranger2
		}

		n := num{num1, num2, num3}

		nums = append(nums, n)
	}

	return nums
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
