package main

import "strings"

func longestCommonPrefix(strs []string) string {
	sb := strings.Builder{}
	n := len(strs)

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < n; j++ {
			if len(strs[j]) == i || strs[j][i] != char {
				return sb.String()
			}
		}
		sb.WriteByte(char)
	}

	return sb.String()
}
