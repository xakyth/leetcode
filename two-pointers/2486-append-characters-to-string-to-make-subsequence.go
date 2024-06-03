package main

func appendCharacters(s string, t string) int {
	i, j := 0, 0
	n, m := len(s), len(t)
	for i < n && j < m {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return m - j
}
