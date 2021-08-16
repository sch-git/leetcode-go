package _32

// middle
// 树：从上到下打印二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if nil == root {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if nil == node {
			continue
		}
		res = append(res, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	return res
}
