package main

type DetectSquares struct {
	pointsFreq map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{pointsFreq: map[[2]int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	this.pointsFreq[[2]int{point[0], point[1]}]++
}

func (this *DetectSquares) Count(point []int) int {
	res := 0
	x1, y1 := point[0], point[1]
	for p, freq := range this.pointsFreq {
		if p[0] == x1 && y1 != p[1] {
			diff := y1 - p[1]
			res += freq * this.pointsFreq[[2]int{x1 - diff, y1}] * this.pointsFreq[[2]int{x1 - diff, p[1]}]
			res += freq * this.pointsFreq[[2]int{x1 + diff, y1}] * this.pointsFreq[[2]int{x1 + diff, p[1]}]
		}
	}

	return res
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
