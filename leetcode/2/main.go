package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0,
		&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
			&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
				&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
					&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
						&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
							&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0,
								&ListNode{1, nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}
	l2 := &ListNode{5,
		&ListNode{6,
			&ListNode{4,
				nil}}}
	l3 := addTwoNumbers(l1, l2)

	for l3 != nil {
		fmt.Print(l3.Val)

		l3 = l3.Next
	}
}
