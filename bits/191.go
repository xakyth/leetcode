package main

func hammingWeight(num int) int {
	cnt := 0
	for num > 0 {
		cnt += int(num & 1)
		num = num >> 1
	}

	return cnt
}
