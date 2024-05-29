package main

func reverseBits(num uint32) uint32 {
	var res uint32 = 0

	j := 31
	for i := 0; i <= 31 && num > 0; i, j = i+1, j-1 {
		x := num & 1
		if x > 0 {
			res += 1 << j
		}
		num = num >> 1
	}

	return res
}
