package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	next := []*TreeNode{}
	if root != nil {
		next = append(next, root)
	}

	for len(next) > 0 {
		tempNext := []*TreeNode{}
		for i := 0; i < len(next); i++ {
			node := next[i]
			if i == len(next)-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				tempNext = append(tempNext, node.Left)
			}
			if node.Right != nil {
				tempNext = append(tempNext, node.Right)
			}
		}
		next = tempNext
	}

	return res
}
