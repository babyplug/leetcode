package range_sum_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RecursiveRangeSumBST(root *TreeNode, l int, r int) int {
	if root == nil {
		return 0
	}
	if root.Val >= l && root.Val <= r {
		return root.Val + RecursiveRangeSumBST(root.Left, l, r) + RecursiveRangeSumBST(root.Right, l, r)
	} else if l < root.Val {
		return RecursiveRangeSumBST(root.Left, l, r)
	} else {
		return RecursiveRangeSumBST(root.Right, l, r)
	}
}

func IterativeRangeSumBST(root *TreeNode, l int, r int) int {
	ans := 0
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			if l <= node.Val && node.Val <= r {
				ans += node.Val
			}
			if l < node.Val { // if l < node.Val we can say that node.Left may be <= l
				stack = append(stack, node.Left)
			}
			if r > node.Val { // if r > node.Val we can say that node.Right may be >= r
				stack = append(stack, node.Right)
			}
		}
	}

	return ans
}
