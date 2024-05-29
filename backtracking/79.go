package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	res := false

	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	var backtrack func(r, c, idx int)
	backtrack = func(r, c, idx int) {
		if idx == len(word)-1 {
			res = true
			return
		}
		if c-1 >= 0 && !visited[r][c-1] && board[r][c-1] == word[idx+1] {
			visited[r][c-1] = true
			backtrack(r, c-1, idx+1)
			visited[r][c-1] = false
		}
		if r-1 >= 0 && !visited[r-1][c] && board[r-1][c] == word[idx+1] {
			visited[r-1][c] = true
			backtrack(r-1, c, idx+1)
			visited[r-1][c] = false
		}
		if c+1 < n && !visited[r][c+1] && board[r][c+1] == word[idx+1] {
			visited[r][c+1] = true
			backtrack(r, c+1, idx+1)
			visited[r][c+1] = false
		}
		if r+1 < m && !visited[r+1][c] && board[r+1][c] == word[idx+1] {
			visited[r+1][c] = true
			backtrack(r+1, c, idx+1)
			visited[r+1][c] = false
		}
	}
	for r := range m {
		for c := range n {
			if board[r][c] == word[0] {
				visited[r][c] = true
				backtrack(r, c, 0)
				visited[r][c] = false
			}
		}
	}

	return res
}
