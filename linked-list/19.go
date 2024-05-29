package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var f func(curr *ListNode) int
	f = func(curr *ListNode) int {
		if curr == nil {
			return 0
		}
		cnt := f(curr.Next)
		if cnt == n {
			curr.Next = curr.Next.Next
			return -1_000_000_000
		}
		return cnt + 1
	}
	cnt := f(head)
	if cnt > 0 {
		head = head.Next
	}
	return head
}
