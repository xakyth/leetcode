package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node.Val == 0 {
			return false
		} else if node.Val == 1 {
			return true
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if node.Val == 2 {
			return l || r
		} else {
			return l && r
		}
	}
	return dfs(root)
}
