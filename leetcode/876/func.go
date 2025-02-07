package main

// ListNode  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var curr, mid int
	midPtr := head
	tail := head
	for tail != nil {
		if (curr+1)/2 > mid {
			midPtr = midPtr.Next
			mid++
		}

		curr++
		tail = tail.Next
	}

	return midPtr
}
