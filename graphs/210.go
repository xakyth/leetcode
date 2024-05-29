package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
	deg := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		deg[edge[0]]++
	}
	queue := make([]int, 0)
	for v, d := range deg {
		if d == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		v := queue[0]
		res = append(res, v)
		queue = queue[1:]
		for _, v2 := range graph[v] {
			deg[v2]--
			if deg[v2] == 0 {
				queue = append(queue, v2)
			}
		}
	}
	if len(res) < numCourses {
		return []int{}
	}
	return res
}
