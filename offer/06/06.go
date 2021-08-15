package _06

// simple

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	temp := make([]int, 0)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}

	res := make([]int, 0)
	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}

	return res
}
