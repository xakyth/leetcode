package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	max := removeNodes(head.Next)
	if head.Next.Val < max.Val {
		head.Next = head.Next.Next
	}
	if head.Val >= max.Val {
		max = head
	}
	return max
}
