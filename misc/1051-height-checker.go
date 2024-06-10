package main

import (
	"sort"
)

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	cnt := 0
	for i, v := range sorted {
		if v != heights[i] {
			cnt++
		}
	}

	return cnt
}
