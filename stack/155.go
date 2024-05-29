package main

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	res := MinStack{}
	res.stack = []int{}
	res.mins = []int{}
	return res
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.stack) == 1 {
		this.mins = append(this.mins, 0)
	} else {
		idx := len(this.stack) - 1
		if val <= this.stack[this.mins[idx-1]] {
			this.mins = append(this.mins, idx)
		} else {
			this.mins = append(this.mins, this.mins[idx-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.mins[len(this.stack)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
