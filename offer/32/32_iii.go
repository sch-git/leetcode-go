package _32

// middle
// 树：从上到下按之字型打印 1：左-右 2：右-左 3：左右
// 实现：双向队列

func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ltor := true
	for len(queue) > 0 {
		temp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			if ltor {
				node := queue[0]
				if nil != node {
					temp = append(temp, node.Val)
					queue = append(queue, node.Left, node.Right)
				}
				queue = queue[1:]
			} else {
				node := queue[len(queue)-1-i]
				if nil != node {
					temp = append(temp, node.Val)
					queue = append([]*TreeNode{node.Right, node.Left}, queue...)
				}
				queue = queue[:len(queue)-1]
			}
		}
		ltor = !ltor
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}

	return res
}
