package main

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	return res
}
