package main

func sortColors(nums []int) {
	count := []int{0, 0, 0}
	for _, v := range nums {
		count[v]++
	}
	for i, k := 0, 0; i < 3; i++ {
		for j := 0; j < count[i]; j, k = j+1, k+1 {
			nums[k] = i
		}
	}
}
