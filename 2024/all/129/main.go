package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sumNumbers(root *TreeNode) int {
	ans := 0
	for _, num := range sum(root) {
		tmp, _ := strconv.Atoi(num)
		ans += tmp
	}
	return ans
}

func sum(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	ls := sum(root.Left)
	rs := sum(root.Right)

	if ls == nil && rs == nil {
		ans := []string{strconv.FormatInt(int64(root.Val), 10)}
		return ans
	}

	ans := make([]string, 0)
	for _, v := range ls {
		ans = append(ans, strconv.FormatInt(int64(root.Val), 10)+v)
	}
	for _, v := range rs {
		ans = append(ans, strconv.FormatInt(int64(root.Val), 10)+v)
	}
	return ans
}
