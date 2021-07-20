package _45

// simple

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	var pre *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || pre == node.Right {
			res = append(res, node.Val)
			pre = node
			root = nil
		} else {
			stack = append(stack, node)
			root = node.Right
		}
	}

	return res
}
