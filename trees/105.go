package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	var dfs func(preStart, inStart, len int) *TreeNode
	dfs = func(preStart, inStart, len int) *TreeNode {
		if len == 0 || preStart >= n {
			return nil
		}
		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		idx := inStart
		for ; inorder[idx] != rootVal; idx++ {
		}

		root.Left = dfs(preStart+1, inStart, idx-inStart)
		root.Right = dfs(preStart+(idx-inStart)+1, idx+1, len-(idx-inStart)-1)

		return root
	}
	return dfs(0, 0, n)
}
