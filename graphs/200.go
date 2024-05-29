package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r == m || c < 0 || c == n || grid[r][c] == '0' || visited[r][c] {
			return
		}
		visited[r][c] = true
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}
	res := 0
	for r := range m {
		for c := range n {
			if !visited[r][c] && grid[r][c] == '1' {
				res++
				dfs(r, c)
			}
		}
	}
	return res
}
