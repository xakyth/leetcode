package main

func canJump(nums []int) bool {
	r := len(nums) - 1

	for i := r - 1; i >= 0; i-- {
		if nums[i] >= r-i {
			r = i
		}
	}
	return r == 0
}
