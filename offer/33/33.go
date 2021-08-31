package _33

import "math"

// middle
// 二叉树/单调栈
// 后续遍历倒序

func verifyPostorder(postorder []int) bool {
	stack := make([]int, 0)
	root := math.MaxInt64
	for i := len(postorder) - 1; i >= 0; i-- {
		stackLen := len(stack)
		if root < postorder[i] {
			return false
		}
		if stackLen == 0 || (stackLen > 0 && postorder[i] > stack[stackLen-1]) {
			stack = append(stack, postorder[i])
			continue
		}
		for stackLen > 0 && postorder[i] < stack[stackLen-1] {
			if root < stack[stackLen-1] {
				return false
			}
			root = stack[stackLen-1]
			stack = stack[:stackLen-1]
			stackLen = len(stack)
		}
		stack = append(stack, postorder[i])
	}
	return true
}
