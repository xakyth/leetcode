package main

func missingNumber(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	n := len(nums)
	return n*(n+1)/2 - sum
}
