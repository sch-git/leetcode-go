package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	q := make([]*TreeNode, 0)
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	q = append(q, root)
	for len(q) > 0 {
		curLen := len(q)
		levelAns := 0
		for ; curLen > 0; curLen-- {
			curNode := q[0]
			q = q[1:]
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
			levelAns = curNode.Val
		}

		ans = append(ans, levelAns)
	}
	return ans
}
