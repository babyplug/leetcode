package shuffle_array

import (
	"fmt"
)

var (
	test1  = []int{2, 5, 1, 3, 4, 7}
	test1N = 3
)

func Solved() {
	fmt.Println("Example 1: ")
	fmt.Println("Input: ", test1)
	fmt.Println("Output: ", shuffle(test1, test1N))
}

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))
	for target := 0; target < n; target++ {
		res = append(res[:target*2], nums[:n][target], nums[n:][target])
	}

	return res
}
