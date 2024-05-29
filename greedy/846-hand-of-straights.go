package main

import (
	"sort"
)

type entry struct {
	val int
	cnt int
}

func isNStraightHand(hand []int, groupSize int) bool {
	hm := make(map[int]int)
	for _, v := range hand {
		hm[v]++
	}
	arr := []entry{}
	for val, cnt := range hm {
		arr = append(arr, entry{val, cnt})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})
	for i := 0; i < len(arr); i++ {
		if arr[i].cnt == 0 {
			continue
		}
		c := arr[i].cnt
		arr[i].cnt = 0
		for j := i + 1; j < i+groupSize; j++ {
			if j >= len(arr) || arr[j-1].val+1 != arr[j].val || arr[j].cnt < c {
				return false
			}
			arr[j].cnt -= c
		}
	}

	return true
}
