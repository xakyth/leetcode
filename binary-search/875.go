package main

import (
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	k := 10_000
	l, r := 1, 10_000
	for l <= r {
		mid := (r-l)/2 + l
		temp := h
		enough := true
		for _, p := range piles {
			temp -= int(math.Ceil(float64(p) / (float64(mid))))
			if temp < 0 {
				enough = false
				break
			}
		}
		if enough {
			r = mid - 1
			k = min(mid, k)
		} else {
			l = mid + 1
		}
	}
	return k
}
