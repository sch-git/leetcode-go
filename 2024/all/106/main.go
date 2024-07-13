package main

func main() {
	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{
			Val: postorder[0],
		}
	}

	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	var leftLen int
	for _, val := range inorder {
		if val == root.Val {
			break
		}
		leftLen++
	}
	root.Left = buildTree(inorder[:leftLen], postorder[:leftLen])
	root.Right = buildTree(inorder[leftLen+1:], postorder[leftLen:len(postorder)-1])
	return root
}
