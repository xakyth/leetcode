package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := -1 << 61
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		v, temp := node.Val, node.Val+l
		temp = max(v+r, temp)
		v = max(temp, v)

		res = max(v, res)
		res = max(node.Val+l+r, res)
		return v
	}
	dfs(root)
	return res
}
