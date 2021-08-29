package _55

// simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	return findDepth(root)
}

func findDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	return max(findDepth(root.Left), findDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
