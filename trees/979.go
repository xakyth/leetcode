package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l != 0 {
			if l < 0 {
				res += -l
			} else {
				res += l
			}
		}
		if r != 0 {
			if r < 0 {
				res += -r
			} else {
				res += r
			}
		}
		return node.Val + l + r - 1
	}
	dfs(root)
	return res
}
