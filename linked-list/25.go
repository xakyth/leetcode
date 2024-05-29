package main

import "fmt"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	h, t, n, c := reverse1(head, 1, k)
	if c < k {
		h, t, _, _ = reverse1(h, 1, c)
		n = nil
	}
	var res *ListNode = nil
	if n != nil {
		res = reverseKGroup(n, k)
	} else {
		t.Next = nil
	}
	t.Next = res

	return h
}

func reverse1(head *ListNode, cnt, k int) (*ListNode, *ListNode, *ListNode, int) {
	if head.Next == nil || cnt == k {
		return head, head, head.Next, cnt
	}
	h, t, n, cnt := reverse1(head.Next, cnt+1, k)
	head.Next.Next = head
	t = head
	return h, t, n, cnt
}

func PrintLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
