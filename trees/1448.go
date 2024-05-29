package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, maxVal int)
	dfs = func(node *TreeNode, maxVal int) {
		if node == nil {
			return
		}
		if node.Val >= maxVal {
			res++
			dfs(node.Left, node.Val)
			dfs(node.Right, node.Val)
		} else {
			dfs(node.Left, maxVal)
			dfs(node.Right, maxVal)
		}
	}
	dfs(root, -(1 << 61))
	return res
}
