package Trie

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{next: make(map[string]*Node)}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

// 获取Trie中存储的单词数量
func (t *Trie) GetSize() int {
	return t.size
}

// 向Trie中添加一个新的单词word
func (t *Trie) Add(word string) {
	cur := t.root
	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = NewNode()
		}
		cur = cur.next[c]
	}

	if cur.isWord == false {
		cur.isWord = true
		t.size++
	}
}
