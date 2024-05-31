package main

func romanToInt(s string) int {
	res := 0
	n := len(s)

	for i := 0; i < n; i++ {
		switch s[i] {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i+1 < n && s[i+1] == 'D' {
				i++
				res += 400
			} else if i+1 < n && s[i+1] == 'M' {
				i++
				res += 900
			} else {
				res += 100
			}
		case 'L':
			res += 50
		case 'X':
			if i+1 < n && s[i+1] == 'C' {
				i++
				res += 90
			} else if i+1 < n && s[i+1] == 'L' {
				i++
				res += 40
			} else {
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i+1 < n && s[i+1] == 'V' {
				i++
				res += 4
			} else if i+1 < n && s[i+1] == 'X' {
				i++
				res += 9
			} else {
				res += 1
			}
		}
	}

	return res
}
