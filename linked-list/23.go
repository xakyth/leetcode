package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		for i := 0; i < len(lists)-1; i += 2 {
			lists[i/2] = mergeTwo(lists[i], lists[i+1])
		}
		if len(lists)%2 == 1 {
			lists[len(lists)/2] = lists[len(lists)-1]
			lists = lists[:len(lists)/2+1]
		} else {
			lists = lists[:len(lists)/2]
		}
	}
	return lists[0]
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	if l1.Val <= l2.Val {
		head, tail, l1 = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next, l1 = l1, l1.Next
			tail = tail.Next
		} else {
			tail.Next, l2 = l2, l2.Next
			tail = tail.Next
		}
	}
	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return head
}
