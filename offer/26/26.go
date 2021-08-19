package _26

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// middle

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, A)
	for len(queue) > 0 {
		node := queue[0]
		if node == nil {
			queue = queue[1:]
			continue
		}

		if node.Val == B.Val && checkNode(node, B) {
			return true
		}

		queue = append(queue[1:], node.Left, node.Right)
	}

	return false
}

func checkNode(a *TreeNode, b *TreeNode) bool {
	// b 为空，a 不管是否为空不影响 b 为 a 的子结构
	if nil == b {
		return true
	}
	// b 非空，a 若为空则说明 b 不是 a 的子结构
	if a == nil {
		return false
	}


	return a.Val == b.Val && checkNode(a.Left, b.Left) && checkNode(a.Right, b.Right)
}
