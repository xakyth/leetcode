package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	cur := []int{}
	sort.Ints(nums)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			res = append(res, append([]int{}, cur...))
			return
		}
		idx := i + 1
		for ; idx < n && nums[idx] == nums[idx-1]; idx++ {
		}
		for j := 0; j < idx-i; j++ {
			cur = append(cur, nums[i])
			backtrack(idx)
		}
		cur = cur[:len(cur)-(idx-i)]
		backtrack(idx)
	}
	backtrack(0)
	return res
}
