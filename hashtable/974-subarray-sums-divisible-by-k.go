package main

func subarraysDivByK(nums []int, k int) int {
	hs := map[int][]int{0: {-1}}
	res := 0

	total := 0
	for i, n := range nums {
		total += n
		rem := total % k
		if rem < 0 {
			rem = k + rem
		}
		indices, contains := hs[rem]
		if !contains {
			hs[rem] = []int{i}
		} else {
			res += len(indices)
			hs[rem] = append(hs[rem], i)
		}
	}
	return res
}
