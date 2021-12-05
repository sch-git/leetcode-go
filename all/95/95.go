package _95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// middle
// 不同的二叉搜索树 II
// 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
func generateTrees(n int) []*TreeNode {
	ops := make([]int, 0)
	for i := 0; i < n; i++ {
		ops = append(ops, i+1)
	}

	res := build(1, n)
	return res
}

func build(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftTrees := build(start, i-1)
		rightTrees := build(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				cur := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				allTrees = append(allTrees, cur)
			}
		}
	}

	return allTrees
}
