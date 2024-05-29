package main

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	res := 0
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || grid[r][c] == 0 {
			return 0
		}
		sum := 1
		visited[r][c] = true
		sum += dfs(r+1, c)
		sum += dfs(r-1, c)
		sum += dfs(r, c-1)
		sum += dfs(r, c+1)
		return sum
	}
	for r := range m {
		for c := range n {
			if !visited[r][c] && grid[r][c] == 1 {
				res = max(dfs(r, c), res)
			}
		}
	}

	return res
}
