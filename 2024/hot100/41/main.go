package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0)
	if root == nil {
		return 0
	}

	inorder(root, &nums)
	if k <= len(nums) {
		return nums[k-1]
	}
	return 0
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}
