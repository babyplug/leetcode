package running_sum

import (
	"testing"
)

func TestRunningSumShouldDisplayLastIndexWithVal10(t *testing.T) {
	result := runningSum([]int{1, 2, 3, 4})
	if result[len(result)-1] != 10 {
		t.Errorf("running sum should display last index with value: 10 but got %v", result)
	}
}

func TestRunningSumShouldDisplayLastIndexWithVal5(t *testing.T) {
	result := runningSum([]int{1, 1, 1, 1, 1})
	if result[len(result)-1] != 5 {
		t.Errorf("running sum should display last index with value: 5 but got %v", result)
	}
}

func TestRunningSumShouldDisplayLastIndexWithVal17(t *testing.T) {
	result := runningSum([]int{3, 1, 2, 10, 1})
	if result[len(result)-1] != 17 {
		t.Errorf("running sum should display last index with value: 17 but got %v", result)
	}
}
