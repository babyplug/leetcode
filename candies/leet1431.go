package candies

import (
	"fmt"
	"math"
)

var (
	case1Candies      = []int{2, 3, 5, 1, 3}
	case1ExtraCandies = 3

	case2Candies      = []int{4, 2, 1, 1, 2}
	case2ExtraCandies = 1

	case3Candies      = []int{12, 1, 12}
	case3ExtraCandies = 10
)

func Solved() {

	fmt.Println("Example 1: ")
	fmt.Println("Input: ", case1Candies, case1ExtraCandies)
	fmt.Println("Output: ", kidsWithCandies(case1Candies, case1ExtraCandies))

	fmt.Println("Example 2: ")
	fmt.Println("Input: ", case2Candies, case2ExtraCandies)
	fmt.Println("Output: ", kidsWithCandies(case2Candies, case2ExtraCandies))

	fmt.Println("Example 3: ")
	fmt.Println("Input: ", case3Candies, case3ExtraCandies)
	fmt.Println("Output: ", kidsWithCandies(case3Candies, case3ExtraCandies))
}

func findMaxNumber(nums []int) int {
	maxValue := math.MinInt64
	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxValue := findMaxNumber(candies)

	for i, kids := range candies {
		res[i] = kids+extraCandies >= maxValue
	}

	return res
}
