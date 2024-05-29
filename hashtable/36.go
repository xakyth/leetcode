package main

func isValidSudoku(board [][]byte) bool {

	var noDuplicates func(group []byte) bool
	noDuplicates = func(group []byte) bool {
		for i, v := range group {
			for j := i + 1; j < 9; j++ {
				if v != '.' && group[j] == v {
					return false
				}
			}
		}
		return true
	}
	for _, row := range board {
		isValid := noDuplicates(row)
		if !isValid {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		temp := make([]byte, 9)
		for j := 0; j < 9; j++ {
			temp[j] = board[j][i]
		}
		isValid := noDuplicates(temp)
		if !isValid {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			temp := make([]byte, 9)
			for y := 0; y < 3; y++ {
				for x := 0; x < 3; x++ {
					temp[y*3+x] = board[i*3+y][j*3+x]
				}
			}
			if !noDuplicates(temp) {
				return false
			}
		}
	}

	return true
}
