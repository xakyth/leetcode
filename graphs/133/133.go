package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	allNodes := make([]Node, 101)
	for i := range 101 {
		allNodes[i].Val = i
	}
	visited := make([]bool, 101)

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil || visited[node.Val] {
			return
		}
		visited[node.Val] = true
		clonedNode := allNodes[node.Val]
		for _, neighbor := range node.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, &allNodes[neighbor.Val])
			dfs(neighbor)
		}
		allNodes[node.Val] = clonedNode

	}
	dfs(node)

	return &allNodes[node.Val]
}
