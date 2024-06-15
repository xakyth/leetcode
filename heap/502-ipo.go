package main

import (
	"container/heap"
	"sort"
)

type project struct {
	capital int
	profit  int
}

type projectHeap []project

func (h projectHeap) Len() int { return len(h) }
func (h projectHeap) Less(i, j int) bool {
	return h[i].profit > h[j].profit
}
func (h projectHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *projectHeap) Push(x any)   { *h = append(*h, x.(project)) }
func (h *projectHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projectsSortedByCapital := make([]project, n)
	for i, p := range profits {
		projectsSortedByCapital[i] = project{capital: capital[i], profit: p}
	}
	sort.Slice(projectsSortedByCapital, func(i, j int) bool {
		return projectsSortedByCapital[i].capital <= projectsSortedByCapital[j].capital
	})

	pHeap := projectHeap{}
	i := 0
	for k > 0 {
		for ; i < n && projectsSortedByCapital[i].capital <= w; i++ {
			heap.Push(&pHeap, projectsSortedByCapital[i])
		}
		if pHeap.Len() == 0 {
			k = 0
		} else {
			w += heap.Pop(&pHeap).(project).profit
			k--
		}
	}

	return w
}
