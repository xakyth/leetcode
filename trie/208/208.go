package main

type Node struct {
	next    [26]*Node
	endWord bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i, ch := range word {
		next := node.next[ch-'a']
		if next == nil {
			next = &Node{}
		}
		if i == len(word)-1 {
			next.endWord = true
		}
		node.next[ch-'a'] = next
		node = next
	}
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for i, ch := range word {
		node = node.next[ch-'a']
		if node == nil {
			return false
		}
		if i == len(word)-1 {
			return node.endWord
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		node = node.next[ch-'a']
		if node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
