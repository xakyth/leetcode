package main

func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n-i)
		for j := range n - i {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j, l, r int) int
	dfs = func(i, j, l, r int) int {
		if dp[j-i][i] != -1 {
			return dp[j-i][i]
		}
		sum := 0
		for k := i; k <= j; k++ {
			tempSum := 0
			if k > i {
				tempSum += dfs(i, k-1, l, nums[k])
			}
			if k < j {
				tempSum += dfs(k+1, j, nums[k], r)
			}
			tempSum += l * nums[k] * r
			if tempSum > sum {
				sum = tempSum
			}
		}
		dp[j-i][i] = sum
		return sum
	}
	dfs(0, n-1, 1, 1)
	return dp[n-1][0]
}
