package main

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(candidates)
	cur := []int{}
	var backtrack func(i, sum int)
	backtrack = func(i, sum int) {
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		} else if i >= n {
			return
		}
		v := candidates[i]
		for j := 0; j*v+sum <= target; j++ {
			for k := 0; k < j; k++ {
				cur = append(cur, v)
			}
			backtrack(i+1, sum+j*v)
			cur = cur[:len(cur)-(j)]
		}
	}
	backtrack(0, 0)
	return res
}
