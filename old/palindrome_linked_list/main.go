package main

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var last *ListNode
	curr := head
	for curr != nil {
		r := *curr
		r.Next = last
		last = &r
		curr = curr.Next
	}
	rHead := last
	r := rHead
	curr = head
	for curr != nil {
		if r.Val != curr.Val {
			return false
		}
		r = r.Next
		curr = curr.Next
	}

	return true
}

func main() {

}
