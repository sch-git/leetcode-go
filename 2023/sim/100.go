package sim

type TreeNode struct {
	Val   int64     `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
