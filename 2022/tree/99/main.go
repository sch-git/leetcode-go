package _9

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var prev *TreeNode
var needChange *TreeNode // [1,3,null,null,2] 需要交换两次
func recoverTree(root *TreeNode)  {
	prev,needChange = nil,nil
	traverse(root)
}

func traverse(root *TreeNode)  {
	if root==nil{
		return
	}

	traverse(root.Left)
	if prev!=nil && root.Val < prev.Val{
		swapVal(root,prev)
		if needChange!=nil{
			swapVal(prev,needChange)
		}else{
			needChange = prev
		}
	}
	prev = root
	traverse(root.Right)
}

func swapVal(root,prev *TreeNode)  {
	root.Val,prev.Val = prev.Val,root.Val
}