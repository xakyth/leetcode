package main

import (
	"math/rand"
)

func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	for _, v := range nums {
		numsMap[v]++
	}
	numCountPairs := make([][2]int, len(numsMap))
	idx := 0
	for num, cnt := range numsMap {
		numCountPairs[idx][0] = num
		numCountPairs[idx][1] = cnt
		idx++
	}

	ans := make([]int, k)
	quickSelect(k, 0, len(numCountPairs), &numCountPairs)
	for i := 0; i < k; i++ {
		ans[i] = numCountPairs[i][0]
	}

	return ans
}

func quickSelect(k, l, r int, numPairs *[][2]int) int {
	pivot := rand.Intn(r-l) + l
	(*numPairs)[l], (*numPairs)[pivot] = (*numPairs)[pivot], (*numPairs)[l]

	idx := l + 1
	for i := l + 1; i < r; i++ {
		if (*numPairs)[i][1] >= (*numPairs)[l][1] {
			(*numPairs)[idx], (*numPairs)[i] = (*numPairs)[i], (*numPairs)[idx]
			idx++
		}
	}
	if idx == k {
		return k
	} else if idx < k {
		return quickSelect(k, idx, r, numPairs)
	} else {
		return quickSelect(k, l, idx, numPairs)
	}
}
