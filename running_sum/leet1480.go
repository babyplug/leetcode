package running_sum

import "fmt"

var (
	test1 = []int{1, 2, 3, 4}
	test2 = []int{1, 1, 1, 1, 1}
	test3 = []int{3, 1, 2, 10, 1}
)

func Solved() {
	fmt.Println("Example 1: ")
	fmt.Println("Input: ", test1)
	fmt.Println("Output: ", runningSum(test1))

	fmt.Println("Example 2: ")
	fmt.Println("Input: ", test2)
	fmt.Println("Output: ", runningSum(test2))

	fmt.Println("Example 3: ")
	fmt.Println("Input: ", test3)
	fmt.Println("Output: ", runningSum(test3))
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}
