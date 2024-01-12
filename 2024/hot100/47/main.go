package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int64]int)
	mp[0] = 1
	var sum int
	var dfs func(node *TreeNode, presum int64)
	dfs = func(node *TreeNode, presum int64) {
		if node == nil {
			return
		}

		presum += int64(node.Val)
		if _, ok := mp[presum-int64(targetSum)]; ok {
			sum += mp[presum-int64(targetSum)]
		}
		mp[presum]++
		dfs(node.Left, presum)
		dfs(node.Right, presum)
		mp[presum]--
		return
	}

	dfs(root, 0)
	return sum
}
