package sim

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && diff(deep(root.Left), deep(root.Right)) <= 1
}

func diff(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDeep := deep(root.Left)
	rightDeep := deep(root.Right)
	return maxNum(leftDeep, rightDeep) + 1
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
