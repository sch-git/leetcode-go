package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	maxNum := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return maxNum(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
