package _725

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, 0)
	nodeCount := 0

	for temp := head; temp != nil; temp = temp.Next {
		nodeCount++
	}

	num := nodeCount / k
	numAdd := nodeCount % k
	curCount := 0
	flag := true
	for temp := head; temp != nil; {
		curCount++
		curNode := temp
		if num <= 0 {
			res = append(res, curNode)
			continue
		}

		if flag {
			res = append(res, curNode)
			flag = false
		}

		if numAdd > 0 && curCount == num+1 {
			flag = true
			temp = temp.Next
			curNode.Next = nil
			numAdd--
		} else if numAdd == 0 && curCount >= num {
			flag = true
			temp = temp.Next
			curNode.Next = nil
		} else {
			temp = temp.Next
		}
	}

	for len(res) < k {
		res = append(res, nil)
	}
	return res
}
