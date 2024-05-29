package main

type Node struct {
	next    [26]*Node
	endWord bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root

	for _, ch := range word {
		next := node.next[ch-'a']
		if next == nil {
			next = &Node{}
		}
		node.next[ch-'a'] = next
		node = next
	}
	node.endWord = true
}

func (this *WordDictionary) Search(word string) bool {
	wordLen := len(word)

	var dfs func(idx int, node *Node) bool
	dfs = func(idx int, node *Node) bool {
		if node == nil {
			return false
		} else if idx == wordLen && node.endWord {
			return true
		} else if idx == wordLen {
			return false
		}

		exist := false
		if word[idx] == '.' {
			for _, n := range node.next {
				exist = exist || dfs(idx+1, n)
				if exist {
					break
				}
			}
		} else {
			exist = exist || dfs(idx+1, node.next[word[idx]-'a'])
		}
		return exist
	}

	return dfs(0, this.root)
}

func main() {
	wd := Constructor()
	wd.AddWord("abc")
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
