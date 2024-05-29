package main

func countSubstrings(s string) int {
	res := len(s)

	for i := 0; i < len(s); i++ {
		for l, r := i-1, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			res++
		}
		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			res++
		}
	}

	return res
}
