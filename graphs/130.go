package main

func solve(board [][]byte) {
	n := len(board[0])
	m := len(board)
	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r < 0 || r >= m || c < 0 || c >= n {
			return false
		}
		if visited[r][c] || board[r][c] == 'X' {
			return true
		}
		visited[r][c] = true
		up := dfs(r-1, c)
		down := dfs(r+1, c)
		left := dfs(r, c-1)
		right := dfs(r, c+1)
		return up && down && left && right
	}
	var dfs2 func(r, c int)
	dfs2 = func(r, c int) {
		if r < 0 || c < 0 || r >= m || c >= n || board[r][c] == 'X' {
			return
		}
		board[r][c] = 'X'
		dfs2(r-1, c)
		dfs2(r+1, c)
		dfs2(r, c-1)
		dfs2(r, c+1)
	}
	for i := range m {
		for j := range n {
			if !visited[i][j] && board[i][j] == 'O' {
				surrounded := dfs(i, j)
				if surrounded {
					dfs2(i, j)
				}
			}
		}
	}
}
