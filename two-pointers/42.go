package main

func trap(height []int) int {
	res := 0
	if len(height) == 1 {
		return res
	}
	l, r := 0, len(height)-1
	res += min(height[l], height[r]) * (r - l - 1)
	for i, j := l, r; i < j; {
		if height[l] <= height[r] {
			i++
		} else {
			j--
		}
		if i >= j {
			break
		}
		if i != l {
			res -= min(height[i], height[l])
			if height[i] > height[l] {
				res += (min(height[i], height[r]) - height[l]) * (r - i - 1)
				l = i
			}
		} else if j != r {
			res -= min(height[j], height[r])
			if height[j] > height[r] {
				res += (min(height[j], height[l]) - height[r]) * (j - l - 1)
				r = j
			}
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
