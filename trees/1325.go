package main

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if l && node.Left.Val == target {
			node.Left = nil
		}
		if r && node.Right.Val == target {
			node.Right = nil
		}
		return node.Left == nil && node.Right == nil
	}
	dfs(root)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
