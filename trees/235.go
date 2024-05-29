package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode = nil
	var dfs func(node *TreeNode) (bool, bool)
	dfs = func(node *TreeNode) (bool, bool) {
		if node == nil || res != nil {
			return false, false
		}

		p1, q1 := dfs(node.Left)
		p2, q2 := dfs(node.Right)
		pCur, qCur := p == node, q == node

		if (p1 && q2) || (q1 && p2) || (pCur && (q1 || q2)) || (qCur && (p1 || p2)) {
			res = node
			return false, false
		}

		return p1 || p2 || pCur, q1 || q2 || qCur
	}
	dfs(root)
	return res
}
