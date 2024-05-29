package main

import (
	"unicode"
)

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tCnt, sCnt := [52]int{}, [52]int{}
	l, r := 0, len(t)
	bestL, bestR := 0, 1_000_000
	for _, r := range t {
		if unicode.IsUpper(r) {
			tCnt[r-'A']++
		} else {
			tCnt[r-'a'+26]++
		}
	}
	isValid := valid(tCnt, sCnt)
	for i := 0; !isValid && i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			sCnt[s[i]-'A']++
		} else {
			sCnt[s[i]-'a'+26]++
		}
		isValid = valid(tCnt, sCnt)
		if isValid {
			r = i + 1
		}
	}
	if !isValid {
		return ""
	}
	incLeft(s, &tCnt, &sCnt, &l)
	bestL, bestR = l, r
	for i := r; i < len(s); i++ {
		if unicode.IsUpper(rune(s[i])) {
			sCnt[s[i]-'A']++
		} else {
			sCnt[s[i]-'a'+26]++
		}
		r++
		incLeft(s, &tCnt, &sCnt, &l)
		if r-l < bestR-bestL {
			bestL, bestR = l, r
		}
	}
	return s[bestL:bestR]
}

func incLeft(s string, tCnt, sCnt *[52]int, l *int) {
	for {
		if unicode.IsUpper(rune(s[*l])) {
			if sCnt[s[*l]-'A']-tCnt[s[*l]-'A'] >= 1 {
				sCnt[s[*l]-'A']--
				*l++
			} else {
				break
			}
		} else {
			if sCnt[s[*l]-'a'+26]-tCnt[s[*l]-'a'+26] >= 1 {
				sCnt[s[*l]-'a'+26]--
				*l++
			} else {
				break
			}
		}
	}
}

func valid(tCnt, sCnt [52]int) bool {
	for i := range tCnt {
		if sCnt[i] < tCnt[i] {
			return false
		}
	}
	return true
}
