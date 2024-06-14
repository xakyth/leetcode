package main

func minIncrementForUnique(nums []int) int {
	n := 100_001
	count := make([]int, n)
	for _, v := range nums {
		count[v]++
	}
	toDistribute := 0
	res := 0

	for _, v := range count {
		if toDistribute == 0 && v <= 1 {
			continue
		}
		toDistribute += v - 1
		res += toDistribute
	}

	if toDistribute > 0 {
		toDistribute--
		res += (toDistribute * (toDistribute + 1)) / 2
	}

	return res
}
