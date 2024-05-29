package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(candidates)

	sort.Ints(candidates)
	cur := []int{}
	var backtrack func(idx, sum int)
	backtrack = func(idx, sum int) {
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		} else if sum > target || idx == n {
			return
		}
		for i := idx; i < n; i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(i+1, sum+candidates[i])
			cur = cur[:len(cur)-1]
		}

	}
	backtrack(0, 0)
	return res
}
