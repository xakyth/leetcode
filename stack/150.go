package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []string{}
	for _, token := range tokens {
		if token == "+" {
			i1, _ := strconv.Atoi(stack[len(stack)-2])
			i2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = fmt.Sprint(i1 + i2)
		} else if token == "-" {
			i1, _ := strconv.Atoi(stack[len(stack)-2])
			i2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = fmt.Sprint(i1 - i2)
		} else if token == "/" {
			i1, _ := strconv.Atoi(stack[len(stack)-2])
			i2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = fmt.Sprint(i1 / i2)
		} else if token == "*" {
			i1, _ := strconv.Atoi(stack[len(stack)-2])
			i2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = fmt.Sprint(i1 * i2)
		} else {
			stack = append(stack, token)
		}
	}
	res, _ := strconv.Atoi(stack[0])
	return res
}
