package main

func singleNumber(nums []int) []int {
	totalXor := 0
	for _, v := range nums {
		totalXor = totalXor ^ v
	}

	shifted := 1
	for i := 0; i < 32; i, shifted = i+1, shifted<<1 {
		if totalXor&shifted > 0 {
			break
		}
	}
	num1 := 0
	for _, v := range nums {
		if shifted&v > 0 {
			num1 = num1 ^ v
		}
	}
	num2 := num1 ^ totalXor

	return []int{num1, num2}
}
