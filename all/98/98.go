package _98

import "math"

// middle

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if nil == root {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
