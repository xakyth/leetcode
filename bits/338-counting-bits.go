package main

func countBits(n int) []int {
	var ones func(m int) int
	ones = func(m int) int {
		ans := 0
		for m > 0 {
			if m&1 == 1 {
				ans++
			}
			m = m >> 1
		}
		return ans
	}
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = ones(i)
	}
	return res
}
