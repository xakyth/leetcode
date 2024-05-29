package main

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[n-1][i] = grid[n-1][i]
	}
	maxInt := int(1 << 62)
	for j := n - 2; j >= 0; j-- {
		for i := 0; i < n; i++ {
			best := maxInt
			for k := 0; k < n; k++ {
				if i == k {
					continue
				}
				best = min(dp[j+1][k], best)
			}
			dp[j][i] = best + grid[j][i]
		}
	}

	best := dp[0][0]
	for _, v := range dp[0] {
		best = min(v, best)
	}
	return best
}
