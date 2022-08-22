package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = 0

func minCameraCover(root *TreeNode) int {
	result = 0
	if traversal(root) == 0 {
		return result + 1
	}

	return result
}

func traversal(root *TreeNode) int {
	if root == nil {
		return 2
	}

	left := traversal(root.Left)
	right := traversal(root.Right)

	if left == 2 && right == 2 {
		return 0
	}

	if left == 0 || right == 0 {
		result++
		return 1
	}

	if left == 1 || right == 1 {
		return 2
	}

	return -1
}
