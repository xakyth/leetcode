package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	visited := make([]bool, n)

	adjList := make([][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	var dfs func(vertex int)
	dfs = func(vertex int) {
		if visited[vertex] {
			return
		}
		visited[vertex] = true
		for _, v := range adjList[vertex] {
			dfs(v)
		}
	}

	dfs(source)
	return visited[destination]
}
