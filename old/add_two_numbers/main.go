package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		switch carry {
		case 0:
			return nil
		default:
			return &ListNode{Val: carry}
		}
	}
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}
	r := &ListNode{}
	s := l1.Val + l2.Val + carry
	c := s / 10
	r.Val = s % 10
	r.Next = add(l1.Next, l2.Next, c)
	return r
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func main() {

}
