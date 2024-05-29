package main

import "strings"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	cols := make([]bool, n)
	diag1 := make([]bool, n*2-1)
	diag2 := make([]bool, n*2-1)

	cur := make([]string, 0)
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			res = append(res, append([]string{}, cur...))
			return
		}
		for c := range n {
			if !cols[c] && !diag1[r+c] && !diag2[n-c-1+r] {
				cols[c], diag1[r+c], diag2[n-c-1+r] = true, true, true
				cur = append(cur, strings.Repeat(".", c)+"Q"+strings.Repeat(".", (n-1-c)))
				backtrack(r + 1)
				cur = cur[:len(cur)-1]
				cols[c], diag1[r+c], diag2[n-c-1+r] = false, false, false
			}
		}
	}
	backtrack(0)
	return res
}
