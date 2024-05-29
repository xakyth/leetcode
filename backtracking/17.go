package main

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	n := len(digits)
	hm := map[rune][]rune{}
	hm['2'] = []rune{'a', 'b', 'c'}
	hm['3'] = []rune{'d', 'e', 'f'}
	hm['4'] = []rune{'g', 'h', 'i'}
	hm['5'] = []rune{'j', 'k', 'l'}
	hm['6'] = []rune{'m', 'n', 'o'}
	hm['7'] = []rune{'p', 'q', 'r', 's'}
	hm['8'] = []rune{'t', 'u', 'v'}
	hm['9'] = []rune{'w', 'x', 'y', 'z'}
	var sb string
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == n {
			res = append(res, sb)
			return
		}
		for _, ch := range hm[rune(digits[idx])] {
			sb += string(ch)
			backtrack(idx + 1)
			sb = sb[:len(sb)-1]
		}
	}
	if n > 0 {

		backtrack(0)
	}
	return res
}
