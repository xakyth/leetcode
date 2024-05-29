package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	res := 1
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		tempMax := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > tempMax {
				tempMax = dp[j] + 1
			}
		}
		dp[i] = tempMax
		res = max(res, tempMax)
	}
	return res
}
