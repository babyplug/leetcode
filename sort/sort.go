package sort

import (
	"fmt"
)

func Sort() {
	fmt.Println("before sort: ", []int{4, 3, 2, 10, 12, 1, 5, 6})
	fmt.Println("bubleSort: ", bubbleSort([]int{4, 3, 2, 10, 12, 1, 5, 6}))
	fmt.Println("selectionSort: ", selectionSort([]int{4, 3, 2, 10, 12, 1, 5, 6}))
	fmt.Println("insertionSort: ", insertionSort([]int{4, 3, 2, 10, 12, 1, 5, 6}))
}

func swap(greater *int, lower *int) {
	temp := *greater
	*greater = *lower
	*lower = temp
}

func bubbleSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}

	return nums
}

func selectionSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		swap(&nums[i], &nums[minIndex])
	}

	return nums
}

func insertionSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		value := nums[i]
		searchIndex := i - 1

		for searchIndex >= 0 && nums[searchIndex] > value {
			nums[searchIndex+1] = nums[searchIndex]
			searchIndex = searchIndex - 1
		}

		nums[searchIndex+1] = value
	}

	return nums
}
