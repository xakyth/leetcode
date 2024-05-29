package main

import "fmt"

func characterReplacement(s string, k int) int {
	best := 0
	kLeft := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		char := s[i]
		temp := k
		r := i
		endReached := false
		for j := i + 1; j < len(s); j++ {
			if s[j] == char {
				r++
			} else {
				temp--
				if temp >= 0 {
					r++
				} else {
					kLeft = temp
					break
				}
			}
			if r == len(s)-1 {
				fmt.Println(string(char))
				endReached = true
			}
			kLeft = temp
		}

		if endReached {
			best = max(best, min(i, kLeft)+r-i+1)
		} else {
			best = max(r-i+1, best)
		}

	}
	return best
}
