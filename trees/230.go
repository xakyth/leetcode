package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	res := -1
	var dfs func(node *TreeNode, cnt int) int
	dfs = func(node *TreeNode, cnt int) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left, cnt)
		cur := cnt + l + 1
		if cur == k {
			res = node.Val
		}
		if res != -1 {
			return cur
		}
		return l + 1 + dfs(node.Right, cur)
	}
	dfs(root, 0)
	return res
}
