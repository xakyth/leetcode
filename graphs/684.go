package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	unionFind := make([]int, n+1)
	for i := range n {
		unionFind[i] = i
	}
	var find func(v int) int
	find = func(v int) int {
		if unionFind[v] == v {
			return v
		}
		return find(unionFind[v])
	}
	union := func(v1, v2 int) {
		unionFind[find(v2)] = find(v1)
	}

	for _, edge := range edges {
		p1 := find(edge[0])
		p2 := find(edge[1])
		if p1 == p2 {
			return edge
		}
		union(p1, p2)
	}
	return edges[n-1]
}
