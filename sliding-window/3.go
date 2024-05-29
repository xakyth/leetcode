package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	best := 0
	for i, r := range s {
		if idx, ok := seen[r]; ok {
			if len(seen) > best {
				best = len(seen)
			}
			for k, v := range seen {
				if v < idx {
					delete(seen, k)
				}
			}
		}
		seen[r] = i
	}
	if len(seen) > best {
		return len(seen)
	}
	return best
}
