package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i <= n; i++ {
		if i == n {
			dp[i] = min(dp[i-1], dp[i-2])
		} else {
			dp[i] = cost[i] + min(dp[i-1], dp[i-2])
		}
	}
	return dp[n]
}
