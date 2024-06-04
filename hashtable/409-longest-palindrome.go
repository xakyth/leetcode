package main

func longestPalindrome(s string) int {
	letterCount := make(map[rune]int)
	for _, char := range s {
		letterCount[char]++
	}
	res := 0
	middle := 0
	for _, count := range letterCount {
		res += (count / 2) * 2
		if count%2 == 1 {
			middle = 1
		}
	}
	return res + middle
}
