package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	l, r := 0, rows*cols
	for l < r {
		mid := (r-l)/2 + l
		i := mid / cols
		j := mid % cols
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}
