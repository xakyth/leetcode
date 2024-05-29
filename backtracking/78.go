package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	cur := []int{}
	var backtrack func(depth int)
	backtrack = func(depth int) {
		if depth == n {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		cur = append(cur, nums[depth])
		backtrack(depth + 1)
		cur = cur[:len(cur)-1]
		backtrack(depth + 1)
	}
	backtrack(0)
	return res
}
