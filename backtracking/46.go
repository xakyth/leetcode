package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	cur := []int{}
	used := make([]bool, n)
	var backtrack func(depth int)
	backtrack = func(depth int) {
		if depth == n {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := range n {
			if !used[i] {
				used[i] = true
				cur = append(cur, nums[i])
				backtrack(depth + 1)
				used[i] = false
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return res
}
