package main

// ListNode is the definition for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func lesser(l1 *ListNode, l2 *ListNode) (less *ListNode, other *ListNode) {
	if l1 == nil {
		return l2, l1
	} else if l2 == nil {
		return l1, l2
	}

	if l1.Val < l2.Val {
		less, other = l1, l2
	} else {
		less, other = l2, l1
	}
	return less, other
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head, other := lesser(l1, l2)
	curr := head
	for curr != nil {
		curr.Next, other = lesser(curr.Next, other)
		curr = curr.Next
	}
	return head
}

// we can divide and conquer merge to get O(n*log(k)) with constant space
// func mergeKLists(lists []*ListNode) *ListNode {
// 	for len(lists) > 1 {
// 		for i := range lists {
// 			e
// 		}
// 	}
// }

func main() {}
