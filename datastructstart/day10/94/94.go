package _4

// simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return res
}
