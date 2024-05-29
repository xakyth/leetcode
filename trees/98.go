package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	res := true
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		cur := node.Val
		low, high := cur, cur

		if node.Left != nil {
			leftLowest, leftHighest := dfs(node.Left)
			if leftHighest >= cur {
				res = false
				return 0, 0
			}
			low = min(leftLowest, low)
		}
		if node.Right != nil {
			rightLowest, rightHighest := dfs(node.Right)
			if cur >= rightLowest {
				res = false
				return 0, 0
			}
			high = max(rightHighest, high)
		}
		return low, high
	}
	dfs(root)

	return res
}
