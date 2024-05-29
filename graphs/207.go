package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		degree[edge[1]]++
	}
	queue := []int{}
	for v, d := range degree {
		if d == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		for _, vertex2 := range graph[vertex] {
			degree[vertex2]--
			if degree[vertex2] == 0 {
				queue = append(queue, vertex2)
			}
		}
	}

	for _, d := range degree {
		if d > 0 {
			return false
		}
	}
	return true
}
