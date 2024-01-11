package main

import (
	"math"
)

func main() {
	root := &TreeNode{Val: -3}
	maxPathSum(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var curMax = -math.MaxInt32
	if root == nil {
		return 0
	}
	cal(root, &curMax)
	return curMax
}

func cal(root *TreeNode, curMax *int) int {
	if root == nil {
		return -math.MaxInt32
	}

	maxNum := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	leftMax := cal(root.Left, curMax)
	rightMax := cal(root.Right, curMax)

	*curMax = maxNum(*curMax, root.Val)
	*curMax = maxNum(*curMax, maxNum(leftMax, rightMax))
	*curMax = maxNum(*curMax, root.Val+maxNum(leftMax, rightMax))
	*curMax = maxNum(*curMax, root.Val+leftMax+rightMax)

	return maxNum(maxNum(leftMax, rightMax)+root.Val, root.Val)
}
