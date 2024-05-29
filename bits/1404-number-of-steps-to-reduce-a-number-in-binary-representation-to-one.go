package main

func numSteps(s string) int {
	oneIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneIdx = i
			break
		}
	}
	if oneIdx == -1 {
		return 1
	}
	n := []rune(s[oneIdx:])
	carry := false
	steps := 0
	for len(n) > 1 {
		steps++
		r := len(n) - 1
		if n[r] == '0' {
			if carry {
				carry = false
			} else {
				n = n[:r]
			}
		} else {
			carry = true
			for i := r; carry && i >= 0; i-- {
				if n[i] == '0' {
					carry = false
					n[i] = '1'
				} else {
					n[i] = '0'
				}
			}
		}

	}
	return steps
}
