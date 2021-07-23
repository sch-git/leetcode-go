package _12

// simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return check(root, 0, targetSum)
}

func check(root *TreeNode, num, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val+num == sum {
		return true
	}

	return check(root.Left, num+root.Val, sum) || check(root.Right, num+root.Val, sum)
}
