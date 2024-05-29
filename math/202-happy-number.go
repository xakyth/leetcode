package main

func isHappy(n int) bool {
	seen := map[int]bool{n: true}

	for n > 1 {
		next := 0

		for n > 0 {
			r := n % 10
			next += r * r
			n = n / 10
		}

		n = next
		if n == 1 {
			return true
		} else if seen[n] {
			return false
		}
		seen[n] = true
	}
	return true
}
