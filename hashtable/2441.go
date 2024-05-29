package main

func findMaxK(nums []int) int {
	res := -1
	hm := map[int]bool{}
	for _, v := range nums {
		if hm[-v] {
			res = max(res, max(-v, v))
		}
		hm[v] = true
	}

	return res
}
