package main

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	var l, r float64 = 0, 1

	for l < r {
		mid := (l + r) / 2
		var maxFraction float64 = 0
		numIdx, denIdx, totalSmaller := 0, 0, 0
		j := 1
		for i := 0; i < n-1; i++ {
			for ; j < n && float64(arr[i]) >= mid*float64(arr[j]); j++ {
			}
			totalSmaller += (n - j)
			if j == n {
				break
			}
			var frac float64 = float64(arr[i]) / float64(arr[j])
			if frac > maxFraction {
				numIdx = i
				denIdx = j
				maxFraction = frac
			}
		}
		if totalSmaller == k {
			return []int{arr[numIdx], arr[denIdx]}
		} else if totalSmaller > k {
			r = mid
		} else {
			l = mid
		}
	}

	return []int{}
}
