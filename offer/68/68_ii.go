package _68

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := dfs(root.Left, p, q)
	right := dfs(root.Right, p, q)

	// p, q 在 root 异侧，则 root 就是最近公共祖先节点
	if left != nil && right != nil {
		return root
	}

	// p, q 都在一侧或其中一个在一侧
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}
