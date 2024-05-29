package main

type node struct {
	key  int
	next *node
}

type entry struct {
	val int
	cnt int
}

type LRUCache struct {
	len  int
	cap  int
	head *node
	tail *node
	hm   [10001]entry
}

func Constructor(capacity int) LRUCache {
	return LRUCache{len: 0, cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	e := this.hm[key]
	if e.cnt <= 0 {
		return -1
	}
	e.cnt++
	this.hm[key] = e
	this.tail.next = &node{key: key}
	this.tail = this.tail.next
	return e.val
}

func (this *LRUCache) Put(key int, value int) {
	newNode := &node{key: key}
	if this.head == nil {
		this.head, this.tail = newNode, newNode
	} else {
		this.tail.next = newNode
		this.tail = this.tail.next
	}
	e := this.hm[key]
	if e.cnt <= 0 {
		this.len++
		e.cnt = 0
	}
	e.cnt++
	e.val = value
	this.hm[key] = e

	if this.len > this.cap {
		this.len--
		for this.head != nil {
			e = this.hm[this.head.key]
			e.cnt--
			this.hm[this.head.key] = e
			this.head = this.head.next
			if e.cnt == 0 {
				break
			}
		}
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
