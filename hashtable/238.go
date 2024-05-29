package main

import "math"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	numsCount := [61]int{}

	for _, v := range nums {
		numsCount[v+30]++
	}

	for i, v := range nums {
		total := 1
		for j, count := range numsCount {
			if v == j-30 {
				if count > 1 {
					total *= int(math.Pow(float64(j-30), float64(count-1)))
				}
			} else {
				if count > 0 {
					total *= int(math.Pow(float64(j-30), float64(count)))
				}
			}
		}
		res[i] = total
	}

	return res
}
