package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	n := head.Next
	var traverse func(t *ListNode) *ListNode
	traverse = func(t *ListNode) *ListNode {
		if t.Next == nil {
			if head.Next == t {
				return nil
			}
			head.Next = t
			t.Next = n
			return n
		}
		h := traverse(t.Next)
		if h == nil {
			return nil
		} else if h == t {
			h.Next = nil
			return nil
		} else if h.Next == t {
			t.Next = nil
			return nil
		}
		h.Next, t.Next = t, h.Next
		return t.Next
	}
	traverse(head)
}
