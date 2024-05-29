package main

import (
	"fmt"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	hashWordsMap := make(map[string][]string)
	lettersCount := [26]int{}

	for _, word := range strs {
		for _, r := range word {
			lettersCount[r-'a']++
		}
		hash := getHashAndReset(&lettersCount)
		hashWordsMap[hash] = append(hashWordsMap[hash], word)
	}
	for _, words := range hashWordsMap {
		ans = append(ans, words)
	}
	return ans
}
func getHashAndReset(lettersCount *[26]int) string {
	var sb strings.Builder
	for i, v := range lettersCount {
		sb.WriteString(fmt.Sprintf("%v%c", v, i+'a'))
		lettersCount[i] = 0
	}
	return sb.String()
}
