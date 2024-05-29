package main

type Cell struct {
	r int
	c int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	queue := []Cell{}
	for r := range m {
		for c := range n {
			if grid[r][c] == 2 {
				queue = append(queue, Cell{r: r - 1, c: c})
				queue = append(queue, Cell{r: r, c: c - 1})
				queue = append(queue, Cell{r: r + 1, c: c})
				queue = append(queue, Cell{r: r, c: c + 1})
			}
		}
	}
	res := 0
	for len(queue) > 0 {
		next := []Cell{}
		flag := false
		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			if cell.r < 0 || cell.r >= m || cell.c < 0 || cell.c >= n || grid[cell.r][cell.c] != 1 {
				continue
			}
			flag = true
			grid[cell.r][cell.c] = 2
			next = append(next, Cell{r: cell.r - 1, c: cell.c})
			next = append(next, Cell{r: cell.r + 1, c: cell.c})
			next = append(next, Cell{r: cell.r, c: cell.c - 1})
			next = append(next, Cell{r: cell.r, c: cell.c + 1})
		}
		queue = next
		if flag {
			res++
		}
	}
	for r := range m {
		for c := range n {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return res
}
