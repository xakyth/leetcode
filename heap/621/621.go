package main

import (
	"container/heap"
)

type taskCount struct {
	task  byte
	count int
}

type taskHeap []taskCount

func (h taskHeap) Len() int           { return len(h) }
func (h taskHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h taskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *taskHeap) Push(x any)        { *h = append(*h, x.(taskCount)) }
func (h *taskHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	tasksCount := make([]int, 26)
	for _, t := range tasks {
		tasksCount[t-'A']++
	}
	h := taskHeap{}
	for i, cnt := range tasksCount {
		if cnt > 0 {
			heap.Push(&h, taskCount{task: byte(i), count: cnt})
		}
	}
	res := 0
	for h.Len() > 0 {
		temp := []taskCount{}
		i := 0
		for ; i <= n && h.Len() > 0; i++ {
			tc := heap.Pop(&h).(taskCount)
			tc.count--
			if tc.count > 0 {
				temp = append(temp, tc)
			}
		}
		if len(temp) > 0 {
			res += (n + 1)
		} else {
			res += i
		}
		for _, t := range temp {
			heap.Push(&h, t)
		}
	}
	return res
}
