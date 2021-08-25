package _34

// middle
// 回溯

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if nil == root {
		return nil
	}

	res := make([][]int, 0)
	temp := make([]int, 0)
	preOrder(root,&temp,&res, target)
	return res
}

func preOrder(root *TreeNode, temp *[]int, res *[][]int, target int) {
	if nil == root {
		return
	}

	target -= root.Val
	*temp = append(*temp, root.Val)
	if isLeaf(root) && target == 0 {
		t := append(make([]int, 0), *temp...)
		*res = append(*res, t)
	}
	preOrder(root.Left, temp, res, target)
	preOrder(root.Right, temp, res, target)
	*temp = (*temp)[:len(*temp)-1]
	return
}

func isLeaf(root *TreeNode) bool {
	return root != nil && root.Left == nil && root.Right == nil
}
