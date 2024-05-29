package main

func pacificAtlantic(heights [][]int) [][]int {
	res := make([][]int, 0)

	n := len(heights[0])
	m := len(heights)
	pac := make([][]bool, m)
	atl := make([][]bool, m)
	for i := range m {
		pac[i] = make([]bool, n)
		atl[i] = make([]bool, n)
	}

	var dfs func(r, c, prev int, reachable [][]bool)
	dfs = func(r, c, prev int, reachable [][]bool) {
		if r < 0 || r >= m || c < 0 || c >= n || heights[r][c] < prev || reachable[r][c] {
			return
		}
		cur := heights[r][c]
		reachable[r][c] = true
		dfs(r, c-1, cur, reachable)
		dfs(r, c+1, cur, reachable)
		dfs(r-1, c, cur, reachable)
		dfs(r+1, c, cur, reachable)
	}
	for i := range m {
		dfs(i, 0, -1, pac)
		dfs(i, n-1, -1, atl)
	}
	for i := range n {
		dfs(0, i, -1, pac)
		dfs(m-1, i, -1, atl)
	}

	for i := range m {
		for j := range n {
			if pac[i][j] && atl[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
