package main

import "strings"

func reversePrefix(word string, ch byte) string {
	idx := strings.IndexByte(word, ch)
	if idx == -1 {
		return word
	}
	runes := []rune(word)
	for i, j := 0, idx; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
