package main

type bar struct {
	height int
	index  int
}

func largestRectangleArea(heights []int) int {
	largestArea := 0
	stack := []bar{}
	for i := len(heights) - 1; i >= 0; i-- {
		area := heights[i]
		if len(stack) == 0 || stack[len(stack)-1].height < heights[i] {
			stack = append(stack, bar{heights[i], i})
		} else if stack[len(stack)-1].height > heights[i] {
			for len(stack) > 0 && stack[len(stack)-1].height > heights[i] {
				b := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					area = max((stack[len(stack)-1].index-i-1)*b.height, area)
					area = max(heights[i]*(stack[len(stack)-1].index-i), area)
				} else {
					area = max((len(heights)-i-1)*b.height, area)
					area = max((len(heights)-i)*heights[i], area)
				}
				largestArea = max(area, largestArea)
			}
			if len(stack) > 0 && stack[len(stack)-1].height == heights[i] {
				stack[len(stack)-1].index = i
			} else {
				stack = append(stack, bar{heights[i], i})
			}
		} else {
			largestArea = max((stack[len(stack)-1].index-i+1)*heights[i], largestArea)
			stack[len(stack)-1].index = i
		}
	}
	startIdx := 0
	if len(stack) > 0 {
		startIdx = stack[len(stack)-1].index
	}
	for len(stack) > 0 {
		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := b.height
		if len(stack) > 0 {
			area = max((stack[len(stack)-1].index-startIdx)*b.height, area)
		} else {
			area = max((len(heights)-startIdx)*b.height, area)
		}
		largestArea = max(area, largestArea)
	}
	return largestArea
}
