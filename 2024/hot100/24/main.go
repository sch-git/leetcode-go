package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	s := make([]int, 0)
	if head == nil || head.Next == nil {
		return true
	}
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
