package main

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	l, r := 0, len(numbers)-1
	for {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			res[0], res[1] = l+1, r+1
			break
		}
	}
	return res
}
