package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := range n {
		dp[0][i] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[m-1][n-1]
}
