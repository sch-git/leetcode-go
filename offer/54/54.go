package _54

// simple
// 二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	if nil == root {
		return 0
	}

	queue := make([]int, 0)
	pre(root, &queue)
	if len(queue)-k >= 0 {
		return queue[len(queue)-k]
	}
	return 0
}

func pre(root *TreeNode, q *[]int) {
	if nil == root {
		return
	}

	pre(root.Left, q)
	*q = append(*q, root.Val)
	pre(root.Right, q)
	return
}
