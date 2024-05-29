package main

func search704(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
