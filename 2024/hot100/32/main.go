package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	oldPre := &Node{Next: head}
	ans := &Node{}
	newHead := ans
	nodemp := make(map[*Node]*Node)
	for head != nil {
		tmp := &Node{Val: head.Val}
		newHead.Next = tmp

		nodemp[head] = tmp
		head = head.Next
		newHead = newHead.Next
	}

	head = oldPre.Next
	for head != nil {
		if head.Random != nil {
			nodemp[head].Random = nodemp[head.Random]
		}
		head = head.Next
	}

	return ans.Next
}
