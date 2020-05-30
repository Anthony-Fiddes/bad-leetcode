package main

// // ListNode is the definition for a singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// // Returns the index of the least node in the slice
// func least(ll []*ListNode) int {
// 	if len(ll) == 1 {
// 		return 0
// 	}
// 	var first int
// 	for first = 0; first < len(ll)-1; first++ {
// 		if ll[first] != nil {
// 			break
// 		}
// 	}
// 	least := first
// 	for i := first + 1; i < len(ll); i++ {
// 		if ll[i] != nil && ll[i].Val < ll[least].Val {
// 			least = i
// 		}
// 	}
// 	return least
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	head := lists[least(lists)]
// 	curr := head
// 	for curr != nil {
// 		lists[least(lists)] = lists[least(lists)].Next
// 		curr.Next = lists[least(lists)]
// 		curr = curr.Next
// 	}
// 	return head
// }

// func main() {}
