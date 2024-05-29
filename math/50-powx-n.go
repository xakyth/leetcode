package main

import "fmt"

func myPow(x float64, n int) float64 {
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	var res float64 = 1
	for n > 0 {
		a := 1
		b := x
		for 2*a <= n {
			b *= b
			a *= 2
		}
		res *= b
		n -= a
	}

	if negative {
		return 1 / res
	} else {
		return res
	}
}

func main() {
	res := myPow(2, 10)

	fmt.Println(res)
}
