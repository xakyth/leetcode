package main

func sumOfDistancesInTree(n int, edges [][]int) []int {

	graph := make([][]int, n)
	dp := make(map[[2]int]int, 2*n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var dfs func(v, parent, level int, cnt *int)
	dfs = func(v, parent, level int, cnt *int) {
		*cnt += level
		for _, v1 := range graph[v] {
			if v1 != parent {
				dfs(v1, v, level+1, cnt)
			}
		}
	}
	var dfs2 func(v, parent int) int
	dfs2 = func(v, parent int) int {
		sum := 1
		for _, v1 := range graph[v] {
			if v1 != parent {
				if dp[[2]int{v, v1}] != 0 {
					sum += dp[[2]int{v, v1}]
				} else {
					temp := dfs2(v1, v)
					dp[[2]int{v, v1}] = temp
					dp[[2]int{v1, v}] = n - temp
					sum += temp
				}

			}
		}
		return sum
	}

	cnt := 0
	for _, v := range graph[0] {
		dfs(v, 0, 1, &cnt)
	}
	res := make([]int, n)
	res[0] = cnt
	queue := []int{}
	queue = append(queue, graph[0]...)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if res[v] != 0 {
			continue
		}
		var visited int
		for _, v1 := range graph[v] {
			if res[v1] != 0 {
				visited = v1
				break
			}
		}
		if dp[[2]int{v, visited}] != 0 {
			cnt = dp[[2]int{v, visited}]
		} else {
			cnt = dfs2(visited, v)
			dp[[2]int{v, visited}] = cnt
			dp[[2]int{visited, v}] = n - cnt
		}
		for _, v1 := range graph[v] {
			if res[v1] == 0 {
				queue = append(queue, v1)
			}
		}
		res[v] = res[visited] + 2*cnt - n
	}

	return res
}
