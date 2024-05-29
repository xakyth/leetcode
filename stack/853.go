package main

import (
	"cmp"
	"slices"
)

type fleet struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	arr := make([]fleet, len(position))
	for i := range position {
		arr[i] = fleet{position[i], float64(target-position[i]) / float64(speed[i])}
	}
	slices.SortFunc(arr, func(a, b fleet) int {
		return cmp.Compare(a.pos, b.pos)
	})
	stack := []fleet{}
	for i := len(arr) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, arr[i])
			continue
		}
		if arr[i].time > stack[len(stack)-1].time {
			stack = append(stack, arr[i])
		}
	}

	return len(stack)
}
