package main

func minOperations(nums []int, k int) int {
	a := k
	for _, n := range nums {
		a = a ^ n
	}
	count := 0
	for a > 0 {
		count += a & 1
		a >>= 1
	}
	return count
}
