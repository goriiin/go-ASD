package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var newHead, tail *ListNode

	for head != nil {
		if head.Val == val {
			head = head.Next

			continue
		}

		if newHead == nil {
			newHead = &ListNode{Val: head.Val}
			tail = newHead
		} else {
			if tail != nil {
				tail.Next = &ListNode{Val: head.Val}
				tail = tail.Next
			}
		}

		head = head.Next
	}

	return newHead
}
