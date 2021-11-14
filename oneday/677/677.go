package _677

func main() {

}

type TreeNode struct {
	key   rune
	value int
	next  []*TreeNode
}

type MapSum struct {
	root *TreeNode
}

func Constructor() MapSum {
	return MapSum{
		root: &TreeNode{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	root := this.root
	for _, v := range key {
		if len(root.next) < 1 {
			root.next = make([]*TreeNode, 0)
			root.next = append(root.next, &TreeNode{
				key: v,
			})
			root = root.next[0]
		} else {
			exist := false
			for _, node := range root.next {
				if node.key == v {
					root = node
					exist = true
				}
			}

			if exist{
				continue
			}
			root.next = append(root.next, &TreeNode{
				key: v,
			})
			root = root.next[len(root.next)-1]
		}
	}
	root.value = val
}

func (this *MapSum) Sum(prefix string) int {
	flag := false
	root := this.root
	for _, pre := range prefix {
		flag = false
		for _, node := range root.next {
			if node.key == pre {
				root = node
				flag = true
				break
			}
		}
		if !flag {
			return 0
		}
	}
	if flag {
		return count(root)
	}

	return 0
}

func count(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if len(node.next) < 1 {
		return node.value
	}

	sum := 0
	for _, child := range node.next {
		sum += count(child)
	}

	return sum + node.value
}
