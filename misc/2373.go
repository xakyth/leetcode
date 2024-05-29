package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	maxLocal := 0
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			maxLocal = 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					maxLocal = max(grid[i+k][j+l], maxLocal)
				}
			}
			res[i] = append(res[i], maxLocal)
		}
	}
	return res
}
