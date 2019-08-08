package TrieR

import "context"

// TrieR 是 Trie in Recursion
// TrieR 将使用递归的方式，实现 Trie 的基本功能
type Node struct {
	isWord bool
	next   map[string]*Node
}

func NewNode() *Node {
	return &Node{next: make(map[string]*Node)}
}

type TrieR struct {
	root *Node
	size int
}

func NewTrieR() *TrieR {
	return &TrieR{root: NewNode()}
}

// 获得Trie中存储的单词数量
func (t *TrieR) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *TrieR) Add(word string) {
	t.add(t.root, word, 0)
}

// 向以node为根的Trie中添加word[index...end)，递归算法
func (t *TrieR) add(node *Node, word string, index int) {
	if index == len(word) {
		if !node.isWord {
			node.isWord = true
			t.size++
		}
		return
	}

	c := string([]rune(word)[index])
	if node.next[c] == nil {
		node.next[c] = NewNode()
	}
	t.add(node.next[c], word, index+1)
}

// 查询单词word是否在Trie中
func (t *TrieR) Contains(word string) bool {
	return contains(t.root, word, 0)
}

// 在以node为根的Trie中查询单词word[index...end)是否存在，递归算法
func contains(node *Node, word string, index int) bool {
	if index == len(word) {
		return node.isWord
	}

	c := string([]rune(word)[index])
	if node.next[c] == nil {
		return false
	}
	return contains(node.next[c], word, index+1)
}

// 查询是否在Trie中有单词以prefix为前缀
func (t *TrieR) IsPrefix(prefix string) bool {
	return isPrefix(t.root, prefix, 0)
}

// 查询在以Node为根的Trie中是否有单词以prefix[index...end)为前缀，递归算法
func isPrefix(node *Node, prefix string, index int) bool {
	if index == len(prefix) {
		return true
	}

	c := string([]rune(prefix)[index])
	if node.next[c] == nil {
		return false
	}

	return isPrefix(node.next[c], prefix, index+1)
}

// 删除word, 返回是否删除成功，递归算法
func (t *TrieR) Remove(word string) bool {
	if word == "" {
		return false
	}
	return t.remove(t.root, word, 0)
}

// 在以Node为根的Trie中删除单词word[index...end),返回是否删除成功，递归算法
func (t *TrieR) remove(node *Node, word string, index int) bool {
	if index == len(word) {
		if !node.isWord {
			return false
		}
		node.isWord = false
		t.size--
		return true
	}

	c := string([]rune(word)[index])
	if node.next[c] == nil {
		return false
	}
	ret := t.remove(node.next[c], word, index+1)
	nextNode := node.next[c]
	if !nextNode.isWord && len(nextNode.next) == 0 {
		node.next.Remove(c)
	}
	return ret
}
