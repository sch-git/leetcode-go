package _05

type TreeNode struct {
	Val   int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	
	root := &TreeNode{Val: preorder[0]}
	index := findIndex(inorder,preorder[0])
	root.Left = buildTree(preorder[1:index+1],inorder[:index])
	root.Right = buildTree(preorder[index+1:],inorder[index+1:])
	return root
}

func findIndex(inorder []int, num int) int {
	for idx, val := range inorder{
		if val == num{
			return idx
		}
	}
	return 0
}