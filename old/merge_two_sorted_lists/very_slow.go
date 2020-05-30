package main

// ListNode is the definition for a singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	if l1 == nil {
// 		return l2
// 	} else if l2 == nil {
// 		return l1
// 	}

// 	var lesser, other *ListNode
// 	if l1.Val < l2.Val {
// 		lesser, other = l1, l2
// 	} else {
// 		lesser, other = l2, l1
// 	}
// 	lesser.Next = mergeTwoLists(lesser.Next, other)
// 	return lesser
// }

// func main() {}
