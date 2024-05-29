package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	} else if n == 3 {
		return max(nums[0], max(nums[1], nums[2]))
	}

	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = nums[0]
	dp1[1] = max(nums[1], nums[0])
	dp2[1] = nums[1]
	for i := 2; i < n-1; i++ {
		dp1[i] = max(dp1[i-2]+nums[i], dp1[i-1])
		dp2[i] = max(dp2[i-2]+nums[i], dp2[i-1])
	}
	dp1[n-1] = dp1[n-2]
	dp2[n-1] = max(dp2[n-2], dp2[n-3]+nums[n-1])
	return max(dp1[n-1], dp2[n-1])
}
