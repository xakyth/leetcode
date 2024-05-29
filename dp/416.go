package main

func canPartition(nums []int) bool {
	neededSum := 0
	for _, v := range nums {
		neededSum += v
	}
	if neededSum%2 == 1 {
		return false
	}
	neededSum /= 2
	dp := make([]bool, neededSum+1)

	dp[0] = true
	for i := range nums {
		for j := neededSum - nums[i]; j >= 0; j-- {
			if dp[j] {
				if j+nums[i] <= neededSum {
					dp[j+nums[i]] = true
				}
			}
		}
	}

	return dp[neededSum]
}
