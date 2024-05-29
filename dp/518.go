package main

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins {
		if c <= amount {
			dp[c]++
		}
		for i := range dp {
			if i-c > 0 {
				dp[i] += dp[i-c]
			}
		}
	}
	return dp[amount]
}
