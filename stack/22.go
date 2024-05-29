package main

import (
	"strings"
)

func generateParenthesis(n int) []string {
	res := []string{}
	stack := []string{}
	var backtrack func(op, clo int)
	backtrack = func(op, clo int) {
		if op == n && clo == n {
			res = append(res, strings.Join(stack, ""))
		} else {
			if op < n {
				stack = append(stack, "(")
				backtrack(op+1, clo)
				stack = stack[:len(stack)-1]
			}
			if clo < op && clo < n {
				stack = append(stack, ")")
				backtrack(op, clo+1)
				stack = stack[:len(stack)-1]
			}
		}
	}
	backtrack(0, 0)
	return res
}
