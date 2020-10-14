package good_pairs

import "testing"

func TestPairShouldDisplay4(t *testing.T) {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3})
	expected := 4
	if result != expected {
		t.Errorf("numIdenticalPairs(nums []int) result: got %v but expect %v", result, expected)
	}
}

func TestPairShouldDisplay6(t *testing.T) {
	result := numIdenticalPairs([]int{1, 2, 3, 1, 1, 3, 3})
	expect := 6
	if result != expect {
		t.Errorf("numIdenticalPairs(nums []int) result: got %v but expect %v", result, expect)
	}
}
