package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	preorder(root)
}

func preorder(root *TreeNode) *TreeNode {
	if root.Right == nil && root.Left == nil {
		return root
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return preorder(root.Right)
	}
	if root.Left == nil {
		return preorder(root.Right)
	}

	node := root.Right
	root.Right = root.Left
	root.Left = nil
	preorder(root.Right).Right = node
	return preorder(node)

}
