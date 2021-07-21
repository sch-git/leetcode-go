package _02

// medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	stack := make([]*TreeNode, 0)
	res := make([][]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		length := len(stack)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := stack[0]
			temp = append(temp, node.Val)
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}
