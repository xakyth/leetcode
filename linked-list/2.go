package main

import (
	"math/big"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := getNumber(l1, 0)
	sum = sum.Add(sum, getNumber(l2, 0))
	if sum.Cmp(big.NewInt(0)) == 0 {
		return &ListNode{Val: 0}
	}
	return getList(sum)
}

func getNumber(h *ListNode, depth int64) *big.Int {
	if h == nil {
		return big.NewInt(0)
	}
	x := new(big.Int).Exp(big.NewInt(10), big.NewInt(depth), nil)
	x = x.Mul(x, big.NewInt(int64(h.Val)))
	x = x.Add(x, getNumber(h.Next, depth+1))
	return x
}

func getList(n *big.Int) *ListNode {
	if n.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	r := new(big.Int).Mod(n, big.NewInt(10))
	n = new(big.Int).Div(n, big.NewInt(10))
	return &ListNode{Val: int(r.Int64()), Next: getList(n)}
}
