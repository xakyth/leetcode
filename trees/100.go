package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	res := true
	var dfs func(r1, r2 *TreeNode)
	dfs = func(r1, r2 *TreeNode) {
		if (r1 == nil && r2 != nil) || (r2 == nil && r1 != nil) {
			res = false
			return
		} else if r1 == nil && r2 == nil {
			return
		}
		if r1.Val != r2.Val {
			res = false
			return
		}
		dfs(r1.Left, r2.Left)
		dfs(r1.Right, r2.Right)
	}
	dfs(p, q)
	return res
}
