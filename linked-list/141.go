package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	hm := map[*ListNode]bool{}
	for head != nil {
		if exist := hm[head]; exist {
			return true
		}
		hm[head] = true
		head = head.Next
	}
	return false
}
