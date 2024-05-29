package main

func longestPalindrome(s string) string {
	res, resLen := string(s[0]), 1
	for i := 0; i < len(s); i++ {
		for l, r := i-1, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l+1 > resLen {
				res = s[l : r+1]
				resLen = len(res)
			}
		}
		for l, r := i, i+1; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
			if r-l+1 > resLen {
				res = s[l : r+1]
				resLen = len(res)
			}
		}
	}

	return res
}
