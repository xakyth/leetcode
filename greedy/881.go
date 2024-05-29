package main

import (
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	res := 0
	l, r := 0, len(people)-1
	for l < r {
		if people[r]+people[l] <= limit {
			res++
			r--
			l++
		} else {
			res++
			r--
		}
	}
	if l == r {
		res++
	}
	return res
}
