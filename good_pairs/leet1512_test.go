package good_pairs

import "testing"

func TestPair(t *testing.T) {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	expected := 4
	if result != expected {
		t.Errorf("numIdenticalPairs(nums []int) result: got %v but expect %v", result, expected)
	}

	result2 := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3, 3})
	expected2 := 6
	if result2 != expected2 {
		t.Errorf("numIdenticalPairs(nums []int) result: got %v but expect %v", result2, expected2)
	}
}
