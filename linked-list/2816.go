package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	h, c := f(head)
	if c > 0 {
		h = &ListNode{Val: c, Next: h}
	}
	return h
}

func f(head *ListNode) (*ListNode, int) {
	if head.Next == nil {
		val := head.Val * 2
		carry := 0
		if val >= 10 {
			carry = val / 10
			val = val % 10
		}
		return &ListNode{Val: val}, carry
	}
	p, carry := f(head.Next)
	val := head.Val*2 + carry
	carry = val / 10
	if val >= 10 {
		val = val % 10
	}

	return &ListNode{Val: val, Next: p}, carry
}
