package main

import (
	"cmp"
	"container/heap"
	"slices"
)

type entry struct {
	val int
	idx int
}

type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *entryHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(entry))
}
func (h *entryHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *entryHeap) Peek() any {
	return (*h)[0]
}

func minInterval(intervals [][]int, queries []int) []int {
	res := make([]int, len(queries))

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	sortedQueries := make([]entry, len(queries))
	for i, v := range queries {
		sortedQueries[i] = entry{v, i}
	}
	slices.SortFunc(sortedQueries, func(e1, e2 entry) int {
		return cmp.Compare(e1.val, e2.val)
	})

	h := &entryHeap{}

	i := 0
	for _, item := range sortedQueries {
		for ; i < len(intervals) && intervals[i][0] <= item.val; i++ {
			heap.Push(h, entry{intervals[i][1] - intervals[i][0] + 1, intervals[i][0]})
		}
		peekVal := -1
		for h.Len() > 0 {
			peek := h.Peek().(entry)
			if peek.idx+peek.val <= item.val {
				heap.Pop(h)
				continue
			}
			peekVal = peek.val
			break
		}
		res[item.idx] = peekVal
	}

	return res
}
