package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(beginWord)
	hm := make(map[string]bool)
	for _, word := range wordList {
		hm[word] = true
	}
	hm[beginWord] = false
	dist := make(map[string]int)
	queue := []string{beginWord}
	dist[beginWord] = 1
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		for i := range n {
			temp1 := word[:i]
			temp2 := word[i+1:]
			for j := 'a'; j <= 'z'; j++ {
				temp := temp1 + string(j) + temp2
				if hm[temp] {
					hm[temp] = false
					dist[temp] = dist[word] + 1
					queue = append(queue, temp)
				}
			}
		}
	}
	return dist[endWord]
}
