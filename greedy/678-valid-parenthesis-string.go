package main

func checkValidString(s string) bool {
	stack := []int{}
	stars := 0
	for _, char := range s {
		if char == '(' {
			stack = append(stack, 0)
		} else if char == '*' {
			if len(stack) == 0 {
				stars++
			} else {
				stack[len(stack)-1]++
			}
		} else {
			if len(stack) == 1 {
				stars += stack[0]
				stack = stack[0:0]
			} else if len(stack) > 1 {
				stack[len(stack)-2] += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				stars--
				if stars < 0 {
					return false
				}
			}
		}
	}
	for len(stack) > 0 {
		elem := stack[len(stack)-1] - 1
		if elem < 0 {
			return false
		}
		if len(stack) > 1 {
			stack[len(stack)-2] += elem
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
