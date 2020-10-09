package good_pairs

import (
	"fmt"
)

var (
	test1 = []int{1, 2, 3, 1, 1, 3}
	test2 = []int{1, 1, 1, 1}
	test3 = []int{1, 2, 3}
)

func Solved() {
	fmt.Println("Example 1: ")
	fmt.Println("Input: ", test1)
	fmt.Println("Output: ", numIdenticalPairs(test1))

	fmt.Println("Example 2: ")
	fmt.Println("Input: ", test2)
	fmt.Println("Output: ", numIdenticalPairs(test2))

	fmt.Println("Example 3: ")
	fmt.Println("Input: ", test3)
	fmt.Println("Output: ", numIdenticalPairs(test3))
}

func numIdenticalPairs(nums []int) int {
	pairs := 0
	m := make(map[int]int)

	for _, v := range nums {
		found := 1

		mv, mk := m[v]
		if mk {
			pairs += mv
			found += mv
		}

		m[v] = found
	}

	return pairs
}
