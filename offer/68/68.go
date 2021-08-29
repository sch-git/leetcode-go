package _68

// simple
// 二叉搜索树 TODO: 未提交

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		}

		return root
	}

	return nil
}
