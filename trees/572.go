package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	res := false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Val == subRoot.Val {
			if isSameTree572(root, subRoot) {
				res = true
				return
			}
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func isSameTree572(p *TreeNode, q *TreeNode) bool {
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
