package main

import (
	"container/heap"
)

type vHeap [][2]int

func (h vHeap) Len() int           { return len(h) }
func (h vHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h vHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *vHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}
func (h *vHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func networkDelayTime(time [][]int, n, k int) int {
	res := -1
	graph := make([][][2]int, n)

	visited := make([]bool, n)
	for _, t := range time {
		graph[t[0]-1] = append(graph[t[0]-1], [2]int{t[1] - 1, t[2]})
	}
	h := &vHeap{[2]int{k - 1, 0}}
	for h.Len() > 0 {
		v := heap.Pop(h).([2]int)
		if visited[v[0]] {
			continue
		}
		visited[v[0]] = true
		if v[0] != k-1 {
			res = max(res, v[1])
		}
		for _, adj := range graph[v[0]] {
			if visited[adj[0]] {
				continue
			}
			heap.Push(h, [2]int{adj[0], adj[1] + v[1]})
		}
	}
	for i := range n {
		if !visited[i] {
			return -1
		}
	}

	return res
}
