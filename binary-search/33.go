package main

func search(nums []int, target int) int {
	k := findShift(nums)
	l, r := k, len(nums)+k
	for l < r {
		m1 := (r-l)/2 + l
		m2 := m1
		if m2 >= len(nums) {
			m2 = m2 - len(nums)
		}
		if nums[m2] == target {
			return m2
		} else if nums[m2] > target {
			r = m1
		} else {
			l = m1 + 1
		}
	}
	return -1
}

func findShift(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return 0
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
