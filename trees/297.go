package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Codec struct {
	sb *strings.Builder
}

func Constructor() Codec {
	return Codec{sb: &strings.Builder{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(node *TreeNode) string {
	if node == nil {
		return ""
	}
	var ser func(node *TreeNode)
	ser = func(node *TreeNode) {
		if node == nil {
			this.sb.WriteRune('n')
			return
		}
		this.sb.WriteString(fmt.Sprintf("(%v", node.Val))
		if node.Left == nil && node.Right == nil {
			this.sb.WriteRune(')')
			return
		}
		ser(node.Left)
		ser(node.Right)
		this.sb.WriteRune(')')
	}
	ser(node)
	return this.sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var deser func(s string, i, n int) (*TreeNode, int)
	deser = func(s string, i, n int) (*TreeNode, int) {
		if i == n {
			return nil, 0
		}
		j := i + 1
		for ; unicode.IsDigit(rune(s[j])) || s[j] == '-'; j++ {
		}
		val, _ := strconv.Atoi(s[i+1 : j])
		newNode := &TreeNode{Val: val}
		var left, right *TreeNode = nil, nil
		nextIdx := -1
		if s[j] == ')' {
			return newNode, j + 1
		} else if s[j] == '(' {
			left, nextIdx = deser(s, j, n)
			if s[nextIdx] == '(' {
				right, nextIdx = deser(s, nextIdx, n)
				nextIdx++
			} else {
				nextIdx += 2
			}
		} else if s[j] == 'n' {
			right, nextIdx = deser(s, j+1, n)
			nextIdx++
		}
		newNode.Left = left
		newNode.Right = right
		return newNode, nextIdx
	}
	p, _ := deser(data, 0, len(data))
	return p
}
