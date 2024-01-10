package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
