package main

func isAnagram(s string, t string) bool {
	letters := make(map[rune]int)
	for _, r := range s {
		letters[r]++
	}
	for _, r := range t {
		letters[r]--
	}
	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}
