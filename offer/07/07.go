package _07

// middle
// 二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	length := 0
	for _, val := range inorder {
		if val == preorder[0] {
			break
		}
		length++
	}

	root.Left = buildTree(preorder[1:length+1], inorder[:length])
	root.Right = buildTree(preorder[length+1:], inorder[length+1:])
	return root
}
