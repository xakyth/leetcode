package main

func getSum(a, b int) int {
	res := 0
	var add func(a, b int) int
	add = func(a, b int) int {
		sum := 0
		c := 0
		i := 0
		for ; (a > 0 || b > 0) && i < 32; i++ {
			n1 := a & 1
			n2 := b & 1
			sum += (n1 ^ n2 ^ c) << i
			if n1&n2 == 1 || n1&c == 1 || n2&c == 1 {
				c = 1
			} else {
				c = 0
			}

			a = a >> 1
			b = b >> 1
		}
		if c > 0 {
			sum += 1 << i
		}
		return sum
	}
	sub := func(a, b int) int {
		res := 0
		c := 0
		for i := 0; a > 0 && i < 32; i++ {
			n1 := a & 1
			n2 := b & 1
			res += (n1 ^ n2 ^ c) << i
			if (n1 == 0 && (n2 > 0 || c > 0)) || (n1 == 1 && (n2 == 1 && c == 1)) {
				c = 1
			} else {
				c = 0
			}
			a = a >> 1
			b = b >> 1
		}
		return res
	}
	if a < 0 {
		a = ^a + 1
		if b < 0 {
			b = ^b + 1
			res = -add(a, b)
		} else {
			if b >= a {
				res = sub(b, a)
			} else {
				res = -sub(a, b)
			}
		}
	} else {
		if b < 0 {
			b = ^b + 1
			if a >= b {
				res = sub(a, b)
			} else {
				res = -sub(b, a)
			}
		} else {
			res = add(a, b)
		}
	}
	return res
}
