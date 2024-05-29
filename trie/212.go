package main

type Node struct {
	next [26]*Node
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}
	var backtrack func(r, c, idx int, prev *Node)
	backtrack = func(r, c, idx int, prev *Node) {
		if r < 0 || r == m || c < 0 || c == n || idx == 10 || visited[r][c] {
			return
		}
		visited[r][c] = true
		node := prev.next[board[r][c]-'a']
		if node == nil {
			node = &Node{}
		}
		prev.next[board[r][c]-'a'] = node
		backtrack(r-1, c, idx+1, node)
		backtrack(r+1, c, idx+1, node)
		backtrack(r, c-1, idx+1, node)
		backtrack(r, c+1, idx+1, node)
		visited[r][c] = false
	}
	root := &Node{}
	for i := range m {
		for j := range n {
			backtrack(i, j, 0, root)
		}
	}
	for _, word := range words {
		node := root
		exist := true
		for _, ch := range word {
			node = node.next[ch-'a']
			if node == nil {
				exist = false
				break
			}
		}
		if exist {
			res = append(res, word)
		}
	}
	return res
}
