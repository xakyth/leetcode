package main

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := nums[0], nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num == 0 {
			curMin, curMax = 0, 0
		} else if num > 0 {
			curMin = min(curMin*num, num)
			curMax = max(curMax*num, num)
		} else {
			curMin, curMax = min(curMax*num, num), max(curMin*num, num)
		}
		res = max(curMax, res)
	}

	return res
}
