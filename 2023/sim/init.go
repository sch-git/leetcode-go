package sim

type TreeNode struct {
	Val   int64     `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}
