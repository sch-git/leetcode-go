package _24

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var ans = math.MinInt32
func maxPathSum(root *TreeNode) int {
	ans = math.MinInt32
	return maxNum(maxPathSum2(root),ans)
}

func maxPathSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxPathSum2(root.Left)
	right := maxPathSum2(root.Right)
	ans = maxNum(ans,left+right+root.Val)
	ans = maxNum(ans,maxNum(maxNum(left,right)+root.Val,root.Val))

	return maxNum(maxNum(left,right)+root.Val,root.Val)
}


func maxNum(x,y int) int{
	if x > y {
		return x
	}
	return y
}