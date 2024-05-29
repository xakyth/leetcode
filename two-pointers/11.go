package main

import "math"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var max float64
	for {
		if l >= r {
			break
		}
		max = math.Max(float64(r-l)*math.Min(float64(height[l]), float64(height[r])), max)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return int(max)
}
