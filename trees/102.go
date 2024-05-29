package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	next := []*TreeNode{root}
	for len(next) > 0 {
		tempNext := []*TreeNode{}
		tempRes := []int{}
		for _, v := range next {
			tempRes = append(tempRes, v.Val)
			if v.Left != nil {
				tempNext = append(tempNext, v.Left)
			}
			if v.Right != nil {
				tempNext = append(tempNext, v.Right)
			}
		}
		if len(tempRes) > 0 {
			res = append(res, tempRes)
		}
		next = tempNext
	}
	return res
}
