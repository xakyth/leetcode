package main

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	slices.SortStableFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		} else {
			if a[1] > b[1] {
				return -1
			} else {
				return 1
			}
		}
	})
	for i := 0; i < len(intervals); i++ {
		l := intervals[i][0]
		r := intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] > r {
				break
			}
			r = max(r, intervals[j][1])
			i++
		}
		res = append(res, []int{l, r})
	}

	return res
}
