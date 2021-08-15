package _35

// middle

// desc: 使用 map 保存原节点和新节点的映射关系，在二次遍历中构建 Random 的关系

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	root := &Node{}
	m := make(map[*Node]*Node)
	r := root
	for cur != nil {
		newNode := &Node{
			Val: cur.Val,
		}
		r.Next = newNode
		m[cur] = newNode
		cur = cur.Next
		r = r.Next
	}

	cur = head
	r = root.Next
	for cur != nil {
		if cur.Random != nil {
			r.Random = m[cur.Random]
		}
		cur = cur.Next
		r = r.Next
	}

	return root.Next
}
