package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		cur := 0
		if s[i-1] != '0' {
			cur += dp[i-1]
		}
		if s[i-2] == '1' {
			cur += dp[i-2]
		} else if s[i-2] == '2' && s[i-1]-'0' <= 6 {
			cur += dp[i-2]
		}
		dp[i] = cur
	}

	return dp[n]
}
