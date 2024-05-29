package main

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry < 10 {
			digits[i] += carry
			return digits
		} else {
			digits[i] = 0
		}
	}
	res := []int{}
	res = append(res, carry)
	res = append(res, digits...)
	return res
}
