package _26

// simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	node := root.Left
	root.Left = root.Right
	root.Right = node
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
