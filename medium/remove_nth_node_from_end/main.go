// https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.2 MB, less than 99.43% of Go online submissions for Remove Nth Node From End of List.
package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Find the length of the linked list
	curr := head
	var length int
	for curr != nil {
		length++
		curr = curr.Next
	}

	// Remove the targeted node
	target := length - n
	if target == 0 {
		return head.Next
	}
	curr = head
	for i := 0; i < length; i++ {
		if i == target-1 {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}

func main() {}
