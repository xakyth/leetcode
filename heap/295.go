package main

import "container/heap"

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{minHeap: MinHeap{}, maxHeap: MaxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(&this.maxHeap, num)
		return
	}
	maxPeek := this.maxHeap.Peek()
	if num > maxPeek {
		heap.Push(&this.minHeap, num)
	} else {
		heap.Push(&this.maxHeap, num)
	}
	if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))
	} else if this.minHeap.Len()-this.maxHeap.Len() > 0 {
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.Peek()+this.minHeap.Peek()) / 2
	} else {
		return float64(this.maxHeap.Peek())
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
