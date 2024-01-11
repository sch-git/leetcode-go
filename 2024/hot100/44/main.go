package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	cur := preorder[0]
	leftNums := make([]int, 0)

	root := &TreeNode{Val: cur}
	if len(preorder) == 1 {
		return root
	}
	for _, val := range inorder {
		if val == cur {
			break
		}
		leftNums = append(leftNums, val)
	}
	rightNum := make([]int, 0)
	if len(leftNums)+1 < len(inorder) {
		rightNum = inorder[len(leftNums)+1:]
	}

	root.Left = buildTree(preorder[1:1+len(leftNums)], leftNums)
	root.Right = buildTree(preorder[len(leftNums):], rightNum)
	return root
}
