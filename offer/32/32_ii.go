package _32

// simple
// 树：从上到下打印，每层一行

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if nil == root {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue[0]
			if nil != node {
				temp = append(temp, node.Val)
				queue = append(queue, node.Left, node.Right)
			}

			queue = queue[1:]
		}
		if len(temp)>0 {
			res = append(res, temp)
		}
	}
	return res
}
