package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeIdxMap := map[*Node]int{}
	c := head
	i := 0
	for ; c != nil; i++ {
		nodeIdxMap[c] = i
		c = c.Next
	}
	nodes := make([]*Node, i+1)
	c = head
	for c != nil {
		idxC := nodeIdxMap[c]
		if nodes[idxC] == nil {
			nodes[idxC] = &Node{Val: c.Val}
		}
		if c.Next != nil {
			idxN := nodeIdxMap[c.Next]
			if nodes[idxN] == nil {
				nodes[idxN] = &Node{Val: c.Next.Val}
			}
			nodes[idxC].Next = nodes[idxN]
		}
		if c.Random != nil {
			idxR := nodeIdxMap[c.Random]
			if nodes[idxR] == nil {
				nodes[idxR] = &Node{Val: c.Random.Val}
			}
			nodes[idxC].Random = nodes[idxR]
		}
		c = c.Next
	}

	return nodes[0]
}
