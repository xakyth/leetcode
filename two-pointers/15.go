package main

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))

	prev := nums[0]
	for i, v := range nums {
		if i > 0 && prev == v {
			continue
		}
		prev = v
		for l, r := i+1, len(nums)-1; l < r; {
			if nums[l]+nums[r]+v == 0 {
				res = append(res, []int{nums[l], nums[r], v})
				l++
				r--
				for {
					if l >= r || nums[l] != nums[l-1] || nums[r] != nums[r+1] {
						break
					}
					l++
					r--
				}
			} else if nums[l]+nums[r]+v > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
