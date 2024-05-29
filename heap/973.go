package main

import (
	"container/heap"
	"math"
)

type pointsHeap [][]int

func (h pointsHeap) Len() int { return len(h) }
func (h pointsHeap) Less(i, j int) bool {
	return math.Sqrt(float64(h[i][0]*h[i][0]+h[i][1]*h[i][1])) < math.Sqrt(float64(h[j][0]*h[j][0]+h[j][1]*h[j][1]))
}
func (h pointsHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *pointsHeap) Push(x any)   { *h = append(*h, x.([]int)) }
func (h *pointsHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	res := make([][]int, k)
	h := pointsHeap(points)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&h).([]int)
	}
	return res
}
