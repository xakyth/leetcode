package main

import "strings"

type node struct {
	endWord  bool
	children [26]*node
}

type trie struct {
	root *node
}

func replaceWords(dictionary []string, sentence string) string {
	tr := trie{root: &node{}}
	for _, word := range dictionary {
		cur := tr.root
		for _, char := range word {
			next := cur.children[char-'a']
			if next == nil {
				next = &node{}
				cur.children[char-'a'] = next
			}
			cur = next
		}
		cur.endWord = true
	}

	sb := strings.Builder{}
	for i := 0; i < len(sentence); i++ {
		cur := tr.root
		foundWord := false
		for ; i < len(sentence); i++ {
			char := rune(sentence[i])
			if char == ' ' {
				sb.WriteRune(char)
				break
			}
			if cur != nil {
				cur = cur.children[char-'a']
				if cur != nil && !foundWord && cur.endWord {
					sb.WriteRune(char)
					foundWord = true
				}
			}
			if !foundWord {
				sb.WriteRune(char)
			}
		}
	}

	return sb.String()
}
