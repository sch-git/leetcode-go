package _27

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// simple

func mirrorTree(root *TreeNode) *TreeNode {
	swapTree(root)
	return root
}

func swapTree(root *TreeNode) {
	if nil == root {
		return
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	swapTree(root.Left)
	swapTree(root.Right)
}
