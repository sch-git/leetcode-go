package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	cal(root, &sum)
	return sum
}

func cal(root *TreeNode, curMax *int) int {
	if root == nil {
		return 0
	}

	left := cal(root.Left, curMax)
	right := cal(root.Right, curMax)

	maxNum := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if left+right > *curMax {
		*curMax = left + right
	}
	return maxNum(left, right) + 1
}
