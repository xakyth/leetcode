package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToUpper(s)

	for l, r := 0, len(s)-1; l < r; {
		if !(unicode.IsDigit(rune(s[l])) || unicode.IsLetter(rune(s[l]))) {
			l++
			continue
		}

		if !(unicode.IsDigit(rune(s[r])) || unicode.IsLetter(rune(s[r]))) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
