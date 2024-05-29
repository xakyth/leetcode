package main

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var backtrack func(x, y, sum int)
	backtrack = func(x, y, sum int) {
		if (x-1 < 0 || grid[y][x-1] == 0) &&
			(y-1 < 0 || grid[y-1][x] == 0) &&
			(x+1 >= n || grid[y][x+1] == 0) &&
			(y+1 >= m || grid[y+1][x] == 0) {
			sum += grid[y][x]
			res = max(sum, res)
			return
		}
		if x-1 >= 0 && grid[y][x-1] > 0 {
			temp := grid[y][x]
			grid[y][x] = 0
			backtrack(x-1, y, sum+temp)
			grid[y][x] = temp
		}
		if y-1 >= 0 && grid[y-1][x] > 0 {
			temp := grid[y][x]
			grid[y][x] = 0
			backtrack(x, y-1, sum+temp)
			grid[y][x] = temp
		}
		if x+1 < n && grid[y][x+1] > 0 {
			temp := grid[y][x]
			grid[y][x] = 0
			backtrack(x+1, y, sum+temp)
			grid[y][x] = temp
		}
		if y+1 < m && grid[y+1][x] > 0 {
			temp := grid[y][x]
			grid[y][x] = 0
			backtrack(x, y+1, sum+temp)
			grid[y][x] = temp
		}
	}
	for i := range m {
		for j := range n {
			if grid[i][j] > 0 {
				backtrack(i, j, 0)
			}
		}
	}

	return res
}
