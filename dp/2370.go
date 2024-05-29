package main

func longestIdealString(s string, k int) int {
	cnt := [26]int{}
	dp := make([]int, len(s)+1)

	for i := 0; i < len(s); i++ {
		r := s[i]
		best := 0
		idx := int(r - 'a')
		for j := idx; j >= 0 && j >= idx-k; j-- {
			best = max(best, cnt[j])
		}
		for j := idx + 1; j < 26 && j <= idx+k; j++ {
			best = max(best, cnt[j])
		}
		cnt[idx] = best + 1
		dp[i+1] = max(cnt[idx], dp[i])
	}
	return dp[len(s)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
