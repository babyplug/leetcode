package range_sum_bst

import (
	"testing"
)

var (
	case1 = TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	case2 = TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
)

func TestCase1ShouldBeDisplay32(t *testing.T) {
	v := RecursiveRangeSumBST(&case1, 7, 15)

	if v != 32 {
		t.Errorf("RecursiveRangeSumBST should be return 32 but got %v", v)
	}
}
func TestCase2ShouldBeDisplay23(t *testing.T) {
	v := RecursiveRangeSumBST(&case2, 6, 10)

	if v != 23 {
		t.Errorf("RecursiveRangeSumBST should be return 23 but got %v", v)
	}
}

func TestCase1IterativeShouldBeDisplay32(t *testing.T) {
	v := IterativeRangeSumBST(&case1, 7, 15)

	if v != 32 {
		t.Errorf("IterativeRangeSumBST should be return 32 but got %v", v)
	}
}
func TestCase2IterativeShouldBeDisplay23(t *testing.T) {
	v := IterativeRangeSumBST(&case2, 6, 11)

	if v != 23 {
		t.Errorf("IterativeRangeSumBST should be return 23 but got %v", v)
	}
}
