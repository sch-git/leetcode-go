package _55

// simple

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}

	return balance(root)>0
}

func balance(root *TreeNode) int {
	if nil == root {
		return 0
	}

	leftHeight := balance(root.Left)
	rightHeight := balance(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}


func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
