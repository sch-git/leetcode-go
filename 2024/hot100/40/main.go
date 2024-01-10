package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root, nil, nil)
}

func check(root, left, right *TreeNode) bool {
	if root == nil {
		return true
	}
	if (left != nil && root.Val <= left.Val) || (right != nil && right.Val <= root.Val) {
		return false
	}

	return check(root.Left, left, root) && check(root.Right, root, right)
}
