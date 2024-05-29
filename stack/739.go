package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for j := len(temperatures) - 1; j >= 0; j-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[j] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[j] = 0
		} else {
			res[j] = stack[len(stack)-1] - j
		}
		stack = append(stack, j)
	}
	return res
}
